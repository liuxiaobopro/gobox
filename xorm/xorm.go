package xorm

import (
	"fmt"
	"html/template"
	"os"
	"strings"

	stringx "github.com/liuxiaobopro/gobox/string"
	"xorm.io/xorm"
)

type GenXormDaoOption func(*genXormDao)

type genXormDao struct {
	Mysql        *xorm.Engine // xorm Engine
	Project      string       // 项目名
	DaoMysqlPath string       // 生成的dao文件绝对路径
	Prefix       string       // 前缀(用于dao文件的生成是否有前缀和package名, 建议传, 没有就不用传了)

	tableInfo []tableInfo // 表信息
}

type tableInfo struct {
	tableName          string   // 表名
	cols               []string // 列名
	programDaoFileName string   // 程序生成的dao文件名
	defaultDaoFileName string   // 用户可修改dao文件名
}

type tplInfo struct {
	Package      string
	PackageUpper string
	Project      string
	Cols         []string
}

var (
	programDaoFileName = "%s_program.go" // 程序生成的dao文件名
	defaultDaoFileName = "%s_default.go" // 用户可修改dao文件名
)

func WithPrefix(prefix string) GenXormDaoOption {
	return func(g *genXormDao) {
		g.Prefix = prefix
	}
}

func NewGenXormDao(mysql *xorm.Engine, daoMysqlPath, projectName string, options ...func(*genXormDao)) *genXormDao {
	if mysql == nil {
		panic("mysql is nil")
	}
	if daoMysqlPath == "" {
		daoMysqlPath = "./dao/mysql"
	}
	if projectName == "" {
		panic("projectName is nil")
	}

	g := &genXormDao{
		Mysql:        mysql,
		DaoMysqlPath: daoMysqlPath,
		Project:      projectName,
	}

	for _, option := range options {
		option(g)
	}
	return g
}

func (g *genXormDao) InitData() error {
	// 获取所有表
	tables, err := g.Mysql.DBMetas()
	if err != nil {
		return err
	}

	// 获取所有表列名
	for _, table := range tables {
		// fmt.Printf("table name: %s\n", table.Name)
		var ti tableInfo
		if g.Prefix != "" {
			// 判断表名前缀是否正确(防止乱传, 不匹配就不截取了)
			if strings.HasPrefix(table.Name, g.Prefix) {
				ti = tableInfo{
					tableName: table.Name[len(g.Prefix):],
				}
			} else {
				ti = tableInfo{
					tableName: table.Name,
				}
			}
		} else {
			ti = tableInfo{
				tableName: table.Name,
			}
		}

		for _, col := range table.Columns() {
			ti.cols = append(ti.cols, col.Name)
			ti.programDaoFileName = fmt.Sprintf(programDaoFileName, ti.tableName)
			ti.defaultDaoFileName = fmt.Sprintf(defaultDaoFileName, ti.tableName)
		}
		g.tableInfo = append(g.tableInfo, ti)
	}
	return nil
}

func (g *genXormDao) Gen() error {
	if err := g.InitData(); err != nil {
		return err
	}

	for _, v := range g.tableInfo {
		// fmt.Println("table name: ", v.tableName)
		programDaoFilePathItem := g.DaoMysqlPath + "/" + v.tableName + "/" + v.programDaoFileName
		defaultDaoFilePathItem := g.DaoMysqlPath + "/" + v.tableName + "/" + v.defaultDaoFileName
		daoPath := g.DaoMysqlPath + "/" + v.tableName
		if _, err := os.Stat(daoPath); err != nil {
			if os.IsNotExist(err) {
				if err := os.MkdirAll(daoPath, os.ModePerm); err != nil {
					return err
				}
			} else {
				return err
			}
		}

		if _, err := os.Stat(programDaoFilePathItem); err != nil {
			if !os.IsNotExist(err) {
				return err
			} else {
				if err := g.createProgramDaoFile(programDaoFilePathItem, v.tableName, v.cols); err != nil {
					return err
				}
			}
		} else {
			// 文件存在，删除
			if err := os.Remove(programDaoFilePathItem); err != nil {
				return err
			}

			// 删除成功，创建
			if err := g.createProgramDaoFile(programDaoFilePathItem, v.tableName, v.cols); err != nil {
				return err
			}
		}

		if _, err := os.Stat(defaultDaoFilePathItem); err != nil {
			if os.IsNotExist(err) {
				// 文件不存在，创建
				if err := g.createDefaultDaoFile(defaultDaoFilePathItem, v.tableName); err != nil {
					return err
				}
			} else {
				return err
			}
		}
	}

	return nil
}

func (g *genXormDao) createProgramDaoFile(path, tName string, cols []string) error {
	var (
		file      *os.File
		err       error
		colsUpper []string
	)

	for _, v := range cols {
		colsUpper = append(colsUpper, stringx.ReplaceCharAfterSpecifiedCharUp(v, "_"))
	}

	if file, err = os.Create(path); err != nil {
		return err
	}
	data := tplInfo{
		Package:      tName,
		PackageUpper: stringx.FirstUp(tName),
		Project:      g.Project,
		Cols:         colsUpper,
	}
	// 解析模板
	tpl, err := template.ParseFiles("xorm/tpl/dao_program.tpl")
	if err != nil {
		return err
	}

	// 应用模板，将结果写入新文件
	err = tpl.Execute(file, data)
	if err != nil {
		return err
	}
	return nil
}

func (g *genXormDao) createDefaultDaoFile(path, tName string) error {
	var (
		file *os.File
		err  error
	)

	if file, err = os.Create(path); err != nil {
		return err
	}
	data := tplInfo{
		Package: tName,
	}

	// 解析模板
	tpl, err := template.ParseFiles("xorm/tpl/dao_default.tpl")
	if err != nil {
		return err
	}

	// 应用模板，将结果写入新文件
	err = tpl.Execute(file, data)
	if err != nil {
		return err
	}
	return nil
}
