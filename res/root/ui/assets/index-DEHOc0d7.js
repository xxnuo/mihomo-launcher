import{s as Ue,a as qe,m as je,t as L,i as m,c,S as $,b as U,d as C,e as V,f as pe,g as _,h as Ae,j as v,k as u,l as Fe,n as ze,r as Ne,o as Re,p as J,u as ge,q as B,v as Be,w as H,x as me,y as He,z as Oe,A as z,B as We,I as Xe,C as K,F as G,D as Z,E as Ve,G as Je,H as Ze,J as Ge,K as Ee,L as Qe,M as Ke,N as Ye,O as et,P as tt,Q as st,R as nt,T as at,U as T,V as ot,W as rt,X as q,Y as lt,Z as it,_ as ct,$ as dt,a0 as ut}from"./vendor-f7bhRhGj.js";(function(){const t=document.createElement("link").relList;if(t&&t.supports&&t.supports("modulepreload"))return;for(const n of document.querySelectorAll('link[rel="modulepreload"]'))s(n);new MutationObserver(n=>{for(const o of n)if(o.type==="childList")for(const r of o.addedNodes)r.tagName==="LINK"&&r.rel==="modulepreload"&&s(r)}).observe(document,{childList:!0,subtree:!0});function a(n){const o={};return n.integrity&&(o.integrity=n.integrity),n.referrerPolicy&&(o.referrerPolicy=n.referrerPolicy),n.crossOrigin==="use-credentials"?o.credentials="include":n.crossOrigin==="anonymous"?o.credentials="omit":o.credentials="same-origin",o}function s(n){if(n.ep)return;n.ep=!0;const o=a(n);fetch(n.href,o)}})();const pt="modulepreload",gt=function(e,t){return new URL(e,t).href},ee={},j=function(t,a,s){let n=Promise.resolve();if(a&&a.length>0){const o=document.getElementsByTagName("link");n=Promise.all(a.map(r=>{if(r=gt(r,s),r in ee)return;ee[r]=!0;const p=r.endsWith(".css"),i=p?'[rel="stylesheet"]':"";if(!!s)for(let g=o.length-1;g>=0;g--){const f=o[g];if(f.href===r&&(!p||f.rel==="stylesheet"))return}else if(document.querySelector(`link[href="${r}"]${i}`))return;const d=document.createElement("link");if(d.rel=p?"stylesheet":pt,p||(d.as="script",d.crossOrigin=""),d.href=r,document.head.appendChild(d),p)return new Promise((g,f)=>{d.addEventListener("load",g),d.addEventListener("error",()=>f(new Error(`Unable to preload CSS for ${r}`)))})}))}return n.then(()=>t()).catch(o=>{const r=new Event("vite:preloadError",{cancelable:!0});if(r.payload=o,window.dispatchEvent(r),!r.defaultPrevented)throw o})},mt=C('<div class="loading loading-spinner">'),ft=C('<button><span class="truncate rounded-none">'),fe=e=>{const[t,a]=Ue(e,["class","loading","icon"]);return(()=>{const s=ft(),n=s.firstChild;return qe(s,je({get class(){return L("btn flex items-center",t.loading?"btn-disabled":t.class)}},a),!1,!0),m(s,c($,{get when(){return t.loading},get children(){return mt()}}),n),m(n,()=>e.icon||e.children),U(()=>n.classList.toggle("flex-1",!t.icon)),s})()},ht=C('<div><div class="collapse-title pr-4 text-xl font-medium after:!top-8"></div><div>'),Ws=e=>{const{title:t,onCollapse:a}=e,s=()=>{const o="collapse-open",r="collapse-close";return e.isOpen?o:r},n=()=>{const o="opacity-100",r="opacity-0";return e.isOpen?o:r};return(()=>{const o=ht(),r=o.firstChild,p=r.nextSibling;return r.$$click=()=>a(!e.isOpen),m(r,t),m(p,c($,{get when(){return e.isOpen},get children(){return pe(()=>e.children)()}})),U(i=>{const l=L(s(),"collapse collapse-arrow select-none border-secondary bg-base-200 shadow-md"),d=L(n(),"collapse-content grid grid-cols-2 gap-2 transition-opacity duration-1000 lg:grid-cols-3 xl:grid-cols-4 2xl:grid-cols-5");return l!==i._v$&&_(o,i._v$=l),d!==i._v$2&&_(p,i._v$2=d),i},{_v$:void 0,_v$2:void 0}),o})()};V(["click"]);const vt=["acid","aqua","autumn","black","bumblebee","business","cmyk","coffee","corporate","cupcake","cyberpunk","dark","dim","dracula","emerald","fantasy","forest","garden","halloween","lemonade","light","lofi","luxury","night","nord","pastel","retro","sunset","synthwave","valentine","winter","wireframe"];var P=(e=>(e.Overview="/overview",e.Proxies="/proxies",e.Rules="/rules",e.Conns="/conns",e.Log="/logs",e.Config="/config",e.Setup="/setup",e))(P||{});const yt=10,Xs={title:{align:"center",style:{color:"gray",fontSize:"16px"}},chart:{toolbar:{show:!1},zoom:{enabled:!1},animations:{easing:"linear"}},noData:{text:"Loading..."},legend:{showForSingleSeries:!0,fontSize:"16px",labels:{colors:"gray"},itemMargin:{horizontal:32}},dataLabels:{enabled:!1},grid:{yaxis:{lines:{show:!1}}},stroke:{curve:"smooth"},tooltip:{enabled:!1},xaxis:{range:yt,labels:{show:!1},axisTicks:{show:!1}},yaxis:{labels:{style:{colors:"gray",fontSize:"13px"},formatter:e=>Ae(e).toString()}}};var Y=(e=>(e[e.NOT_CONNECTED=-1]="NOT_CONNECTED",e[e.MEDIUM=200]="MEDIUM",e[e.HIGH=500]="HIGH",e))(Y||{}),he=(e=>(e[e.NOT_CONNECTED=-1]="NOT_CONNECTED",e[e.MEDIUM=800]="MEDIUM",e[e.HIGH=1500]="HIGH",e))(he||{}),ve=(e=>(e.OFF="off",e.DOTS="dots",e.BAR="bar",e.Auto="auto",e))(ve||{}),I=(e=>(e.NATURAL="orderNatural",e.LATENCY_ASC="orderLatency_asc",e.LATENCY_DESC="orderLatency_desc",e.NAME_ASC="orderName_asc",e.NAME_DESC="orderName_desc",e))(I||{}),O=(e=>(e.EN="en-US",e.ZH="zh-CN",e))(O||{}),ye=(e=>(e.Details="details",e.Close="close",e.ID="ID",e.Type="type",e.Process="process",e.Host="host",e.SniffHost="sniffHost",e.Rule="rules",e.Chains="chains",e.DlSpeed="dlSpeed",e.ULSpeed="ulSpeed",e.Download="dl",e.Upload="ul",e.ConnectTime="connectTime",e.SourceIP="sourceIP",e.SourcePort="sourcePort",e.Destination="destination",e.InboundUser="inboundUser",e))(ye||{});const te=200,be=Object.values(ye),bt={...Object.fromEntries(be.map(e=>[e,!0])),ID:!1};var k=(e=>(e.XS="xs",e.SM="sm",e.MD="md",e.LG="lg",e))(k||{}),wt=(e=>(e.Global="global",e.Rule="rule",e.Direct="direct",e))(wt||{}),we=(e=>(e.Info="info",e.Error="error",e.Warning="warning",e.Debug="debug",e.Silent="silent",e))(we||{});const xt=[200,300,500,800,1e3],Pt=xt[0],Ct={add:"Add",overview:"Overview",proxies:"Proxies",proxiesSettings:"Proxies Settings",rules:"Rules",connections:"Connections",connectionsSettings:"Connections Settings",connectionsDetails:"Connections Details",logs:"Logs",logsSettings:"Logs Settings",config:"Config",upload:"Upload",download:"Download",uploadTotal:"Upload Total",downloadTotal:"Download Total",activeConnections:"Active Connections",memoryUsage:"Memory Usage",traffic:"Traffic",memory:"Memory",down:"Down",up:"Up",proxyProviders:"Proxy Providers",ruleProviders:"Rule Providers",search:"Search",ID:"ID",type:"Type",name:"Name",process:"Process",host:"Host",sniffHost:"Sniff Host",chains:"Chains",connectTime:"Time",dlSpeed:"DL Speed",ulSpeed:"UL Speed",dl:"DL",ul:"UL",sourceIP:"Source IP",sourcePort:"Source Port",destination:"Destination",inboundUser:"Inbound User",close:"Close",reset:"Reset",dnsQuery:"DNS Query",dots:"Dots",bar:"Bar",auto:"Auto",off:"Off",proxiesPreviewType:"Proxies preview type",urlForIPv6SupportTest:"URL for IPv6 support test",urlForLatencyTest:"URL for latency test",autoCloseConns:"Automatically Close connections",useTwemoji:"Use Twemoji Mozilla Font",autoSwitchTheme:"Automatically switch theme",favDayTheme:"Favorite light theme",favNightTheme:"Favorite dark theme",renderInTwoColumns:"Render in two columns",updateGEODatabases:"Update GEO Databases",restartCore:"Restart Core",upgradeCore:"Upgrade Core",proxiesSorting:"Proxies Sorting",orderNatural:"Original order in config file",orderLatency_asc:"By latency from low to high",orderLatency_desc:"By latency from high to low",orderName_asc:"By name alphabetically (A-Z)",orderName_desc:"By name alphabetically (Z-A)",ms:"ms",updated:"Updated",tableSize:"Table size",logLevel:"Log Level",info:"info",silent:"silent",debug:"debug",warning:"warning",error:"error",logMaxRows:"Log Maximum Reserved Rows",xs:"Extra small size",sm:"Small size",md:"Normal size",lg:"Large size",switchEndpoint:"Switch Endpoint",switchLanguage:"Switch Language",latencyTestTimeoutDuration:"Latency Test Timeout Duration",all:"All",sequence:"Sequence",payload:"Payload",details:"Details",endpointURL:"Endpoint URL",secret:"Secret",global:"Global",rule:"Rule",direct:"Direct",active:"Active",closed:"Closed",sort:"Sort",hideUnAvailableProxies:"Hide UnAvailable Proxies",reloadConfig:"Reload Config",flushFakeIP:"Flush Fake-IP",tagClientSourceIPWithName:"Tag Client Source IP With Name",tag:"Tag",coreConfig:"Core Config",xdConfig:"XD Config",version:"Version",expire:"Expire",noExpire:"Null",allowLan:"Allow Lan",enableTunDevice:"Enable TUN Device",tunModeStack:"TUN Mode Stack",tunDeviceName:"TUN Device Name",interfaceName:"Interface Name",en:"English",zh:"Chinese",port:"{{ name }} Port"},$t={add:"添加",overview:"概览",proxies:"代理",proxiesSettings:"代理设置",rules:"规则",connections:"连接",connectionsSettings:"连接设置",connectionsDetails:"连接详情",logs:"日志",logsSettings:"日志设置",config:"配置",upload:"上传",download:"下载",uploadTotal:"上传总量",downloadTotal:"下载总量",activeConnections:"活动连接",memoryUsage:"内存使用情况",traffic:"流量",memory:"内存",down:"下载",up:"上传",proxyProviders:"代理提供者",ruleProviders:"规则提供者",search:"搜索",ID:"ID",type:"类型",name:"名字",process:"进程",host:"主机",sniffHost:"嗅探域名",chains:"链路",connectTime:"连接时间",dlSpeed:"下载速度",ulSpeed:"上传速度",dl:"下载量",ul:"上传量",sourceIP:"源地址",sourcePort:"源端口",destination:"目标地址",inboundUser:"入站用户",close:"关闭",reset:"重置",dnsQuery:"DNS 查询",dots:"点阵",bar:"条形",auto:"自适应",off:"关闭",proxiesPreviewType:"节点组预览样式",urlForIPv6SupportTest:"测试 IPv6 支持链接",urlForLatencyTest:"测速链接",autoCloseConns:"自动断开连接",useTwemoji:"使用 Twemoji Mozilla 字体",autoSwitchTheme:"自动切换主题",favDayTheme:"浅色主题偏好",favNightTheme:"深色主题偏好",renderInTwoColumns:"双列渲染",updateGEODatabases:"更新 GEO 数据库",restartCore:"重启核心",upgradeCore:"更新核心",proxiesSorting:"节点排序",orderNatural:"原配置文件中的排序",orderLatency_asc:"按延迟从低到高",orderLatency_desc:"按延迟从高到低",orderName_asc:"按名称字母排序 (A-Z)",orderName_desc:"按名称字母排序 (Z-A)",ms:"毫秒",updated:"更新于",tableSize:"表格大小",logLevel:"日志等级",info:"信息",silent:"静默",debug:"调试",warning:"警告",error:"错误",logMaxRows:"日志最大保留行数",xs:"超小尺寸",sm:"小尺寸",md:"正常尺寸",lg:"超大尺寸",switchEndpoint:"切换后端",switchLanguage:"切换语言",latencyTestTimeoutDuration:"测速超时时间",all:"全部",sequence:"序列号",payload:"内容",details:"详情",endpointURL:"后端地址",secret:"密钥",global:"全局",rule:"规则",direct:"直连",active:"活动",closed:"已关闭",sort:"排序",hideUnAvailableProxies:"隐藏不可用节点",reloadConfig:"重载配置",flushFakeIP:"清空 Fake-IP",tagClientSourceIPWithName:"为客户端源 IP 地址添加名称标记",tag:"标记",coreConfig:"核心配置",xdConfig:"XD 配置",version:"版本",expire:"到期时间",noExpire:"不限时",allowLan:"允许局域网访问",enableTunDevice:"开启 TUN 转发",tunModeStack:"TUN 模式堆栈",tunDeviceName:"TUN 设备名称",interfaceName:"接口名称",en:"英文",zh:"中文",port:"{{ name }} 端口"},xe={[O.EN]:Ct,[O.ZH]:$t},[St,Vs]=v(u(Reflect.has(xe,navigator.language)?navigator.language:O.EN),{name:"lang",storage:localStorage}),[Tt,Dt]=Fe(e=>[ze(()=>Re(xe[e.locale]),Ne)]),Pe=()=>Dt(),[Js,Zs]=v(u(ve.Auto),{name:"proxiesPreviewType",storage:localStorage}),[Gs,Es]=v(u(I.NATURAL),{name:"proxiesOrderingType",storage:localStorage}),[Qs,Ks]=v(u(!1),{name:"hideUnAvailableProxies",storage:localStorage}),[Ys,en]=v(u(!0),{name:"renderProxiesInTwoColumns",storage:localStorage}),[W,tn]=v(u("https://www.gstatic.com/generate_204"),{name:"urlForLatencyTest",storage:localStorage}),[E,sn]=v(u("https://api-ipv6.ip.sb/ip"),{name:"urlForIPv6SupportTest",storage:localStorage}),[Lt,nn]=v(u(!1),{name:"autoCloseConns",storage:localStorage}),[It,an]=v(u(!0),{name:"useTwemoji",storage:localStorage}),[kt,on]=v(u(!1),{name:"autoSwitchTheme",storage:localStorage}),[Mt,rn]=v(u("nord"),{name:"favDayTheme",storage:localStorage}),[_t,ln]=v(u("sunset"),{name:"favNightTheme",storage:localStorage}),[cn,dn]=v(u(k.XS),{name:"connectionsTableSize",storage:localStorage}),[un,pn]=v(u(bt),{name:"connectionsTableColumnVisibility",storage:localStorage}),[gn,mn]=v(u(be),{name:"connectionsTableColumnOrder",storage:localStorage}),[fn,hn]=v(u([]),{name:"clientSourceIPTags",storage:localStorage}),[vn,yn]=v(u(k.XS),{name:"logsTableSize",storage:localStorage}),[bn,wn]=v(u(we.Info),{name:"logLevel",storage:localStorage}),[xn,Pn]=v(u(Pt),{name:"logMaxRows",storage:localStorage}),Cn=e=>{let t="table-xs";switch(e){case k.XS:t="table-xs";break;case k.SM:t="table-sm";break;case k.MD:t="table-md";break;case k.LG:t="table-lg";break}return t},[R,$n]=v(u(5e3),{name:"latencyTestTimeoutDuration",storage:localStorage}),Ut=()=>W().startsWith("https"),D=()=>Ut()?he:Y,[qt,Ce]=u([]),[$e,jt]=u(null),Sn=()=>{const[e,t]=u([]),[a,s]=u([]),[n,o]=u(!1);return J(()=>{var p;const r=(p=$e())==null?void 0:p.connections;r&&ge(()=>{const i=Se(r,a());if(At(a()),!n()){const l=Ft(i,qt());s(i),t(l.slice(-te))}Ce(l=>l.slice(-(i.length+te)))})}),{closedConnections:e,activeConnections:a,paused:n,setPaused:o}},Se=(e,t)=>{const a=new Map;return t.forEach(s=>a.set(s.id,s)),e.map(s=>{const n=a.get(s.id);return!n||!B.isNumber(n.download)||!B.isNumber(n.upload)?{...s,downloadSpeed:0,uploadSpeed:0}:{...s,downloadSpeed:s.download-n.download,uploadSpeed:s.upload-n.upload}})},At=e=>{Ce(t=>B.unionWith(t,e,(a,s)=>a.id===s.id))},Ft=(e,t)=>B.differenceWith(t,e,(a,s)=>a.id===s.id),Tn=()=>w().delete("connections"),zt=e=>w().delete(`connections/${e}`),[Dn,se]=u(!1),[Ln,ne]=u(!1),[In,ae]=u(!1),[kn,oe]=u(!1),[Mn,re]=u(!1),_n=async()=>{const e=w();se(!0);try{await e.put("configs",{searchParams:{force:!0},json:{path:"",payload:""}})}catch{}se(!1)},Un=async()=>{const e=w();ae(!0);try{await e.post("cache/fakeip/flush")}catch{}ae(!1)},qn=async()=>{const e=w();ne(!0);try{await e.post("configs/geo")}catch{}ne(!1)},jn=async()=>{const e=w();oe(!0);try{await e.post("upgrade")}catch{}oe(!1)},An=async()=>{const e=w();re(!0);try{await e.post("restart")}catch{}re(!1)},Fn=()=>w().get("configs").json(),zn=async(e,t,a)=>{try{await w().patch("configs",{json:{[e]:t}}).json(),await a()}catch(s){Be.error(s.message)}},Nn=async()=>{const e=w(),{version:t}=await e.get("version").json();return t},Nt=()=>w().get("providers/proxies").json(),Rt=()=>w().get("proxies").json(),le=e=>w().put(`providers/proxies/${e}`),Te=e=>w().get(`providers/proxies/${e}/healthcheck`,{timeout:5*1e3}).json(),Bt=(e,t)=>w().put(`proxies/${e}`,{body:JSON.stringify({name:t})}),ie=(e,t,a,s)=>{const n=w();return t!==""?Te(t).then(o=>({delay:o[e]})):n.get(`proxies/${e}/delay`,{searchParams:{url:a,timeout:s}}).json()},ce=(e,t,a)=>w().get(`group/${e}/delay`,{searchParams:{url:t,timeout:a}}).json(),Rn=()=>w().get("rules").json(),Bn=()=>w().get("providers/rules").json(),Hn=e=>w().put(`providers/rules/${e}`),On=async e=>{const t="https://api.github.com/repos/MetaCubeX/mihomo",a=/(alpha|beta|meta)-?(\w+)/.exec(e);if(!a)return!1;const s=a[1],n=a[2];if(s==="meta"){const{assets:o}=await H.get(`${t}/releases/latest`).json();return!o.some(({name:p})=>p.includes(n))}if(s==="alpha"){const{assets:o}=await H.get(`${t}/releases/tags/Prerelease-Alpha`).json();return!o.some(({name:p})=>p.includes(n))}return!1},Wn=e=>e.includes("sing-box"),Xn=e=>/^https?/.test(e)?e:`${window.location.protocol}//${e}`,N=()=>{const[e,t]=u({}),a=(n,o)=>{t({...e(),[n]:o})};return{map:e,set:a,setWithCallback:async(n,o)=>{a(n,!0);try{await o()}catch{}a(n,!1)}}},Ht=(e="")=>{const t=e.toLowerCase();return t.includes("shadowsocks")?t.replace("shadowsocks","ss"):t==="hysteria"?"hy":t==="wireguard"?"wg":t},Ot=(e="")=>{const t=e.toLowerCase();return!["selector","direct","reject","urltest","loadbalance","fallback","relay"].includes(t)},Vn=(e,t,a,s)=>a===I.NATURAL?e:e.sort((n,o)=>{if(s!=null&&s.has(n)&&!(s!=null&&s.has(o)))return-1;if(s!=null&&s.has(o)&&!(s!=null&&s.has(n)))return 1;const r=t[n],p=t[o];switch(a){case I.LATENCY_ASC:return r===D().NOT_CONNECTED?1:p===D().NOT_CONNECTED?-1:r-p;case I.LATENCY_DESC:return r===D().NOT_CONNECTED?1:p===D().NOT_CONNECTED?-1:p-r;case I.NAME_ASC:return n.localeCompare(o);case I.NAME_DESC:return o.localeCompare(n);default:return 0}}),Jn=(e,t,a,s)=>s?e.filter(n=>a!=null&&a.has(n)?!0:t[n]!==D().NOT_CONNECTED):e,{map:Wt,set:Xt}=N(),{map:Vt,setWithCallback:Jt}=N(),{map:Zt,setWithCallback:Gt}=N(),{map:Et,setWithCallback:Qt}=N(),{map:Kt,setWithCallback:Yt}=N(),[es,de]=u(!1),[ts,ss]=u([]),[ns,as]=u(new Set),[ue,os]=u([]),[De,Q]=u({}),[Le,F]=u({}),[Ie,rs]=u({}),ls=e=>{var p;const t={...Ie()},a={...De()},s={...Le()},n=(i,l,d=!0)=>{var f,y,x,b;const g=(f=i.extra)==null?void 0:f[l];if(Array.isArray(g)){const h=(y=g.at(-1))==null?void 0:y.delay;if(h)return h}if(d)return(b=(x=i.history)==null?void 0:x.at(-1))==null?void 0:b.delay},o={};e.forEach(i=>{const{udp:l,xudp:d,type:g,now:f,name:y,provider:x=""}=i;if(t[i.name]={udp:l,xudp:d,type:g,now:f,name:y,provider:x},!f)a[i.name]=n(i,W())||D().NOT_CONNECTED;else if(a[f]!==void 0)a[i.name]=a[f];else{const h=o[f]??new Set;h.add(i.name),o[f]=h}const b=(n(i,E(),!1)??0)>0;s[i.name]=b});const r=Object.keys(o).filter(i=>a[i]!==void 0);for(;r.length>0;){const i=r.shift(),l=a[i];for(const d of((p=o[i])==null?void 0:p.values())??[])a[d]=l,r.push(d)}me(()=>{rs(t),Q(a),F(s)})},X=()=>{const e=async()=>{const[{providers:l},{proxies:d}]=await Promise.all([Nt(),Rt()]),g=[...d.GLOBAL.all??[],"GLOBAL"],f=Object.values(d).filter(b=>{var h;return(h=b.all)==null?void 0:h.length}).sort((b,h)=>g.indexOf(b.name)-g.indexOf(h.name)),y=Object.values(l).filter(b=>b.name!=="default"&&b.vehicleType!=="Compatible"),x=[...Object.values(d),...y.flatMap(b=>b.proxies.filter(h=>!(h.name in d)).map(h=>({...h,provider:b.name})))];me(()=>{ss(f),as(new Set(["DIRECT","REJECT",...f.map(b=>b.name)])),os(y),ls(x)})},t=async(l,d)=>{await Bt(l.name,d),await e(),Lt()&&ge(()=>{var f;const g=Se(((f=$e())==null?void 0:f.connections)??[],[]);g.length>0&&g.forEach(({id:y,chains:x})=>{x.includes(l.name)&&zt(y)})})},a=async(l,d)=>{const g=E();if(!g||g.length===0){F({});return}let f=!1;try{const{delay:y}=await ie(l,d,g,R());f=y>0}catch{f=!1}F(y=>({...y,[l]:f}))},s=async l=>{const d=E();if(!d||d.length===0){F({});return}const g=await ce(l,d,R()),f=Object.fromEntries(Object.entries(g).map(([y,x])=>[y,x>0]));F(y=>({...y,...f}))};return{collapsedMap:Wt,setCollapsedMap:Xt,proxyIPv6SupportMap:Le,proxyLatencyTestingMap:Vt,proxyGroupLatencyTestingMap:Zt,proxyProviderLatencyTestingMap:Et,updatingMap:Kt,isAllProviderUpdating:es,proxies:ts,proxyGroupNames:ns,proxyProviders:ue,proxyLatencyTest:async(l,d)=>{Jt(l,async()=>{const{delay:g}=await ie(l,d,W(),R());Q(f=>({...f,[l]:g}))}),await a(l,d)},proxyGroupLatencyTest:async l=>{Gt(l,async()=>{const d=await ce(l,W(),R());Q(g=>({...g,...d})),await e()}),await s(l)},latencyMap:De,proxyNodeMap:Ie,fetchProxies:e,selectProxyInGroup:t,updateProviderByProviderName:l=>Yt(l,async()=>{try{await le(l)}catch{}await e()}),updateAllProvider:async()=>{de(!0);try{await Promise.allSettled(ue().map(l=>le(l.name))),await e()}finally{de(!1)}},proxyProviderLatencyTest:l=>Qt(l,async()=>{await Te(l),await e()})}},[is,Zn]=v(u(""),{name:"selectedEndpoint",storage:localStorage}),[cs,Gn]=v(u([]),{name:"endpointList",storage:localStorage}),w=()=>{const e=M();if(!e)return H.create({});const t=new Headers;return e.secret&&t.set("Authorization",`Bearer ${e.secret}`),H.create({prefixUrl:e.url,headers:t})},M=()=>cs().find(({id:e})=>e===is()),ds=()=>{var e;return(e=M())==null?void 0:e.secret},us=()=>{var e;return new URL(((e=M())==null?void 0:e.url)??"").origin.replace("http","ws")},ps=(e,t={})=>{const a=new URLSearchParams(t);a.set("token",ds()??"");const s=He(`${us()}/${e}?${a.toString()}`),n=Oe(s,"message");return z(()=>{var r;return n()?JSON.parse((r=n())==null?void 0:r.data):null})},[gs,ke]=v(u("sunset"),{name:"theme",storage:localStorage}),ms=C('<li class="tooltip tooltip-bottom">'),fs=C('<div class="drawer drawer-end w-auto sm:ml-auto"><input id=themes type=checkbox class=drawer-toggle><div class="drawer-content flex items-center"><label for=themes class="btn btn-circle btn-primary drawer-button btn-sm"></label></div><div class="drawer-side overflow-x-hidden"><label for=themes class=drawer-overlay></label><ul class="menu gap-2 rounded-l-box bg-base-300 p-2">'),hs=C('<li class="btn btn-xs">'),vs=C('<div class="navbar-center hidden lg:flex"><ul class="menu menu-horizontal menu-lg gap-2 p-0">'),ys=C('<ul class="navbar z-50 flex w-auto items-center justify-center bg-base-300 px-4 shadow-lg"><div class="navbar-start gap-4"><div class="drawer w-auto lg:hidden"><input id=navs type=checkbox class=drawer-toggle><div class="drawer-content flex items-center"><label for=navs class="btn btn-circle drawer-button btn-sm"></label></div><div class=drawer-side><label for=navs class=drawer-overlay></label><ul class="menu min-h-full w-2/5 gap-2 rounded-r-box bg-base-300 pt-20"></ul></div></div></div><div class=navbar-end><div class="flex items-center gap-2">'),bs=C("<li>"),ws=({href:e,tooltip:t,children:a})=>(()=>{const s=ms();return Z(s,"data-tip",t),m(s,c(K,{class:"rounded-box",href:e,children:a})),s})(),xs=()=>(()=>{const e=fs(),t=e.firstChild,a=t.nextSibling,s=a.firstChild,n=a.nextSibling,o=n.firstChild,r=o.nextSibling;return m(s,c(Ve,{})),m(r,c(G,{each:vt,children:p=>(()=>{const i=hs();return i.$$click=()=>ke(p),Z(i,"data-theme",p),m(i,p),i})()})),e})(),Ps=()=>{const[e]=Pe(),t=()=>[{href:P.Overview,name:e("overview"),icon:c(Je,{})},{href:P.Proxies,name:e("proxies"),icon:c(Ze,{})},{href:P.Rules,name:e("rules"),icon:c(Ge,{})},{href:P.Conns,name:e("connections"),icon:c(Ee,{})},{href:P.Log,name:e("logs"),icon:c(Qe,{})},{href:P.Config,name:e("config"),icon:c(Ke,{})}],a=We(),[s,n]=u(!1);return(()=>{const o=ys(),r=o.firstChild,p=r.firstChild,i=p.firstChild,l=i.nextSibling,d=l.firstChild,g=l.nextSibling,f=g.firstChild,y=f.nextSibling,x=r.nextSibling,b=x.firstChild;return i.addEventListener("change",h=>n(h.target.checked)),m(d,c(Xe,{})),m(y,c(G,{get each(){return t()},children:({href:h,name:A})=>(()=>{const S=bs();return S.$$click=()=>n(!1),m(S,c(K,{href:h,children:A})),S})()})),m(r,c(Ls,{}),null),m(o,c($,{get when(){return a.pathname!==P.Setup},get children(){const h=vs(),A=h.firstChild;return m(A,c(G,{get each(){return t()},children:({href:S,name:Me,icon:_e})=>c(ws,{href:S,tooltip:Me,children:_e})})),h}}),x),m(b,c(xs,{})),U(()=>i.checked=s()),o})()};V(["click"]);const Cs=C('<span class="badge badge-sm p-px"><span class=scale-75>IPv6'),$s=e=>{const{proxyIPv6SupportMap:t}=X(),a=z(()=>t()[e.name]===!0);return c($,{get when(){return a()},get children(){return Cs()}})},Ss=C("<span>"),Ts=e=>{const[t]=Pe(),{latencyMap:a}=X(),[s,n]=u(""),o=z(()=>a()[e.name]);return J(()=>{n("text-success"),o()>D().HIGH?n("text-error"):o()>D().MEDIUM&&n("text-warning")}),c($,{get when(){return z(()=>typeof o()=="number")()&&o()!==Y.NOT_CONNECTED},get children(){const r=Ss();return m(r,o,null),m(r,()=>t("ms"),null),U(()=>_(r,`whitespace-nowrap text-xs ${s()}`)),r}})},Ds=C('<div class="text-md flex items-center gap-1 whitespace-nowrap font-bold uppercase sm:text-xl"><span>(</span><a class="text-primary transition-transform hover:rotate-90 hover:scale-125"href=https://github.com/metacubex/metacubexd target=_blank>xd</a><span>)'),Ls=()=>(()=>{const e=Ds(),t=e.firstChild;return m(e,c(K,{class:"bg-gradient-to-br from-primary to-secondary bg-clip-text text-transparent",get href(){return M()?"/":"/setup"},children:"metacube"}),t),e})(),Is=C('<div class="sticky bottom-0 z-50 flex items-center justify-end bg-base-100 bg-opacity-80 p-4 backdrop-blur"><div class="flex justify-end gap-2">'),ks=C('<dialog class="modal modal-bottom sm:modal-middle"><div class="modal-box p-0"><div><div class="flex items-center gap-4 text-xl font-bold"><span></span></div></div><div class=p-4></div></div><form method=dialog class=modal-backdrop><button>'),Ms="sticky bottom-0 z-50 flex items-center justify-end bg-base-100 bg-opacity-80 p-4 backdrop-blur",En=e=>{let t;return(()=>{const a=ks(),s=a.firstChild,n=s.firstChild,o=n.firstChild,r=o.firstChild,p=n.nextSibling;return Ye(i=>{var l;return(t=i)&&((l=e.ref)==null?void 0:l.call(e,i))},a),s.$$contextmenu=i=>i.preventDefault(),m(o,()=>e.icon,r),m(r,()=>e.title),m(n,c(fe,{class:"btn-circle btn-sm",onClick:()=>t==null?void 0:t.close(),get icon(){return c(et,{size:20})}}),null),m(p,pe(()=>e.children)),m(s,c($,{get when(){return e.action},get children(){const i=Is(),l=i.firstChild;return m(l,()=>e.action),i}}),null),U(()=>_(n,L(Ms,"top-0 justify-between"))),a})()};V(["contextmenu"]);const _s=C('<div><div class="flex items-center justify-between gap-2"><span class="break-all text-left text-sm"></span><span class="flex items-center gap-1"></span></div><div class="flex items-center justify-between gap-1"><div></div><div class=text-xs>'),Qn=e=>{const{proxyLatencyTest:t,proxyLatencyTestingMap:a}=X(),{proxyName:s,isSelected:n,onClick:o}=e,{proxyNodeMap:r}=X(),p=z(()=>r()[s]),i=()=>{var l,d,g;return Ot((l=p())==null?void 0:l.type)?(d=p())!=null&&d.xudp?"xudp":(g=p())!=null&&g.udp?"udp":null:null};return(()=>{const l=_s(),d=l.firstChild,g=d.firstChild,f=g.nextSibling,y=d.nextSibling,x=y.firstChild,b=x.nextSibling;return tt(l,"click",o,!0),Z(l,"title",s),m(g,s),m(f,c($s,{get name(){return e.proxyName}}),null),m(f,c(fe,{class:"btn-circle btn-ghost h-auto min-h-0 w-auto",get icon(){return c(st,{size:20,get class(){return L(a()[s]&&"animate-pulse text-success")}})},onClick:h=>{h.stopPropagation(),t(s,p().provider)}}),null),m(x,()=>{var h;return Ht((h=p())==null?void 0:h.type)},null),m(x,c($,{get when(){return i()},get children(){return` :: ${i()}`}}),null),m(b,c(Ts,{get name(){return e.proxyName}})),U(h=>{const A=L("card card-bordered tooltip-bottom flex flex-col justify-between gap-1 border-neutral-focus bg-neutral p-2 text-neutral-content",n&&"border-primary bg-primary-content text-primary",o&&"cursor-pointer"),S=L("text-xs text-slate-500",n&&"text-primary");return A!==h._v$&&_(l,h._v$=A),S!==h._v$2&&_(x,h._v$2=S),h},{_v$:void 0,_v$2:void 0}),l})()};V(["click"]);const Us=C('<div><div class="flex-1 overflow-y-auto p-2 sm:p-4"><div class=pb-8>'),qs=q(()=>j(()=>import("./Setup-1jhU3Tvl.js"),__vite__mapDeps([0,1]),import.meta.url)),js=q(()=>j(()=>import("./Overview-yOUd8Ylh.js"),__vite__mapDeps([2,1]),import.meta.url)),As=q(()=>j(()=>import("./Connections-hDEkMReA.js"),__vite__mapDeps([3,1,4,5,6]),import.meta.url)),Fs=q(()=>j(()=>import("./Logs-G54ijlcT.js"),__vite__mapDeps([7,1,4,5]),import.meta.url)),zs=q(()=>j(()=>import("./Proxies-ynI51cVX.js"),__vite__mapDeps([8,1,6,5]),import.meta.url)),Ns=q(()=>j(()=>import("./Rules-P3bEEeTK.js"),__vite__mapDeps([9,1,6]),import.meta.url)),Rs=q(()=>j(()=>import("./Config-c8CDVxAf.js"),__vite__mapDeps([10,1,5]),import.meta.url)),Bs=()=>{const e=ps("connections");return J(()=>jt(e())),null},Hs=()=>{const e=nt();return J(()=>{kt()&&ke(e()?_t():Mt())}),c(Tt,{get locale(){return St()},get children(){const t=Us(),a=t.firstChild,s=a.firstChild;return m(t,c(Ps,{}),a),m(s,c(at,{get children(){return[c($,{get when(){return M()},get children(){return[c(T,{get path(){return P.Overview},component:js}),c(T,{get path(){return P.Proxies},component:zs}),c(T,{get path(){return P.Rules},component:Ns}),c(T,{get path(){return P.Conns},component:As}),c(T,{get path(){return P.Log},component:Fs}),c(T,{get path(){return P.Config},component:Rs}),c(T,{path:"*",get element(){return c(ot,{get href(){return P.Overview}})}})]}}),c(T,{get path(){return M()?P.Setup:"*"},component:qs})]}}),null),m(s,c($,{get when(){return M()},get children(){return c(Bs,{})}}),null),m(t,c(rt,{position:"bottom-center"}),null),U(n=>{const o=L("relative flex h-screen flex-col overscroll-y-none subpixel-antialiased",It()?"font-twemoji":"font-no-twemoji"),r=gs();return o!==n._v$&&_(t,n._v$=o),r!==n._v$2&&Z(t,"data-theme",n._v$2=r),n},{_v$:void 0,_v$2:void 0}),t}})};lt.extend(it);ct(()=>c(ut,{get source(){return dt()},get children(){return c(Hs,{})}}),document.getElementById("root"));export{Qs as $,wn as A,fe as B,yt as C,Xs as D,Pn as E,xt as F,vn as G,bn as H,xn as I,nn as J,tn as K,we as L,En as M,$n as N,sn as O,Es as P,I as Q,Ks as R,en as S,k as T,Zs as U,ve as V,Lt as W,W as X,R as Y,E as Z,Gs as _,is as a,Ys as a0,Js as a1,D as a2,Ts as a3,X as a4,Jn as a5,Vn as a6,Ws as a7,Qn as a8,Rn as a9,vt as aA,ln as aB,Mt as aC,_t as aD,It as aE,On as aF,Fn as aG,Bn as aa,Hn as ab,N as ac,St as ad,Nn as ae,Wn as af,w as ag,zn as ah,Dn as ai,_n as aj,Ln as ak,qn as al,In as am,Un as an,kn as ao,jn as ap,Mn as aq,An as ar,P as as,wt as at,O as au,an as av,Vs as aw,on as ax,kt as ay,rn as az,Zn as b,ps as c,M as d,cs as e,be as f,bt as g,dn as h,fn as i,hn as j,cn as k,$e as l,qt as m,Sn as n,ye as o,zt as p,gn as q,un as r,Gn as s,Xn as t,Pe as u,Tn as v,mn as w,pn as x,Cn as y,yn as z};
function __vite__mapDeps(indexes) {
  if (!__vite__mapDeps.viteFileDeps) {
    __vite__mapDeps.viteFileDeps = ["./Setup-1jhU3Tvl.js","./vendor-f7bhRhGj.js","./Overview-yOUd8Ylh.js","./Connections-hDEkMReA.js","./index-NBB5o6lb.js","./ConfigTitle-ER2TWONh.js","./global-k9w1ZtkR.js","./Logs-G54ijlcT.js","./Proxies-ynI51cVX.js","./Rules-P3bEEeTK.js","./Config-c8CDVxAf.js"]
  }
  return indexes.map((i) => __vite__mapDeps.viteFileDeps[i])
}