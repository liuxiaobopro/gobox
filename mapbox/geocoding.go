package mapbox

import (
	"fmt"

	httpx "github.com/liuxiaobopro/gobox/http"
)

type CountryStr string  // 国家代码
type LanguageStr string // 语言代码

const (
	// CountryAbroad 国外
	CountryAbroad = "ad,af,ae,ag,ai,al,am,ao,aq,ar,as,at,au,aw,ax,az,ba,bb,bd,be,bf,bg,bh,bi,bj,bl,bm,bn,bo,bq,br,bs,bt,bv,bw,by,bz,ca,cc,cd,cf,cg,ch,ci,ck,cl,cm,cn,co,cr,cv,cu,cw,cx,cy,cz,de,dj,dk,dm,do,dz,ec,ee,eg,eh,er,es,et,fi,fj,fk,fm,fo,fr,ga,gb,gd,ge,gf,gg,gh,gi,gl,gm,gn,gp,gq,gr,gs,gt,gu,gw,gy,hk,hm,hn,hr,ht,hu,id,ie,il,im,in,io,iq,ir,is,it,je,jm,jo,jp,ke,kg,kh,ki,km,kn,kp,kr,kw,ky,kz,la,lb,lc,lk,li,lr,ls,lt,lu,lv,ly,ma,mc,me,md,mf,mg,mh,mk,ml,mm,mn,mo,mp,mq,mr,ms,mt,mu,mv,mw,mx,my,mz,na,nc,ne,nf,ng,ni,nl,no,np,nr,nu,nz,om,pa,pe,pf,pg,ph,pk,pl,pm,pn,pr,ps,pt,pw,py,qa,re,ro,rs,ru,rw,sa,sb,sc,sd,se,sg,sh,si,sj,sk,sl,sm,sn,so,sr,ss,st,sv,sx,sy,sz,tc,td,tf,tg,th,tj,tk,tl,tm,tn,to,tr,tt,tv,tw,tz,ua,ug,um,us,uy,uz,va,vc,ve,vg,vi,vn,vu,wf,ws,xs,xk,ye,yt,za,zm,zw"

	// CountryChina 国内
	CountryChina = "cn"

	// CountryGlobal 全球
	CountryGlobal = CountryAbroad + "," + CountryChina

	// LanguageChinese 中文
	LanguageChinese = "zh"
)

type Geocoding struct {
	Mapbox

	Country  CountryStr
	Language string

	Q string
}

func WithCountry(country CountryStr) func(*Geocoding) {
	return func(g *Geocoding) {
		g.Country = country
	}
}

func WithLanguage(language string) func(*Geocoding) {
	return func(g *Geocoding) {
		g.Language = language
	}
}

func WithQuery(q string) func(*Geocoding) {
	return func(g *Geocoding) {
		g.Q = q
	}
}

func NewGeocoding(accessToken string, options ...func(*Geocoding)) *Geocoding {
	g := &Geocoding{
		Mapbox: Mapbox{
			AccessToken: accessToken,
		},
		Country:  CountryGlobal,
		Language: LanguageChinese,
		Q:        "",
	}

	for _, option := range options {
		option(g)
	}

	return g
}

func (g *Geocoding) Query() ([]byte, error) {
	client := &httpx.Client{
		Url: g.url(),
	}

	return client.Get()
}

func (g *Geocoding) url() string {
	return fmt.Sprintf("https://api.mapbox.com/geocoding/v5/mapbox.places/%s.json?country=%s&language=%s&access_token=%s", g.Q, g.Country, g.Language, g.AccessToken)
}
