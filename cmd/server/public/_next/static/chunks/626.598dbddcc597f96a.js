(self.webpackChunk_N_E=self.webpackChunk_N_E||[]).push([[626,182,678],{27888:function(e,t,a){"use strict";a.d(t,{Z:function(){return x}});var r=a(85893),n=a(67294),s=a(41664),i=a.n(s),o=a(25675),l=a.n(o);class c{constructor(e,t){this.href=e,this.title=t}}var h=a(91580),p=a.n(h),d=a(77546),g=a(2078),u=a(11163);let v=n.forwardRef(function({title:e,titleId:t,...a},r){return n.createElement("svg",Object.assign({xmlns:"http://www.w3.org/2000/svg",fill:"none",viewBox:"0 0 24 24",strokeWidth:1.5,stroke:"currentColor","aria-hidden":"true",ref:r,"aria-labelledby":t},a),e?n.createElement("title",{id:t},e):null,n.createElement("path",{strokeLinecap:"round",strokeLinejoin:"round",d:"M15.75 9V5.25A2.25 2.25 0 0013.5 3h-6a2.25 2.25 0 00-2.25 2.25v13.5A2.25 2.25 0 007.5 21h6a2.25 2.25 0 002.25-2.25V15M12 9l-3 3m0 0l3 3m-3-3h12.75"}))});var _=a(31730);class f extends _.Z{get pages(){let e=[];return this.props.hideAllPages||(e.push(new c("/agenda","Agenda")),this.props.pages&&e.push(...this.props.pages)),e}render(){return(0,r.jsx)("nav",{className:p().navbar,children:(0,r.jsxs)("div",{className:p().container,children:[(0,r.jsx)("div",{className:p().rightSideContainer,children:(0,r.jsxs)(i(),{href:"/",className:p().logoContainer,children:[(0,r.jsx)(l(),{src:"/logo.png",width:33,height:35,className:p().logo,alt:"Pet Journal Logo"}),(0,r.jsx)("span",{className:p().logoTitle,children:"Pet Journal"})]})}),(0,r.jsxs)("div",{className:p().leftSideContainer,children:[(0,r.jsx)("div",{className:p().pagesContainer,children:this.pages.map(e=>(0,r.jsx)(i(),{href:e.href,className:p().pageContainer,children:(0,r.jsx)("span",{className:"".concat(p().pageLink," ").concat(this.props.router.pathname===e.href?p().pageSelected:""),children:e.title})},e.title))}),!this.props.hideAllPages&&(0,r.jsxs)("div",{className:"flex flex-row gap-2 items-center",children:[(0,r.jsx)(d.Z,{icon:(0,r.jsx)(g.Z,{className:"flex p-1"}),onCLick:()=>this.props.router.push("/account"),className:"hover:bg-indigo-200 hover:text-indigo-700 h-[30px] w-[30px]"}),(0,r.jsx)(v,{className:"h-10 w-10 text-slate-300 p-2 rounded-full hover:bg-gray-600",onClick:()=>this.logout(()=>this.props.router.replace("/auth/login"))})]}),this.props.buttons&&(0,r.jsx)("div",{className:p().buttonsContainer,children:this.props.buttons})]})]})})}constructor(e){super(e),this.state={}}}var x=(0,u.withRouter)(f)},62472:function(e,t,a){"use strict";a.r(t);var r=a(85893);a(67294);var n=a(27888),s=a(81280),i=a(31730),o=a(11163);class l extends i.Z{componentDidMount(){let e=this.cookies.get("token");void 0!==e?(this.setState({token:e}),this.initInterval=setInterval(()=>{this.props.router.isReady&&(this.props.init&&this.props.init(e),clearInterval(this.initInterval))},1)):this.logout(()=>this.props.router.replace("/auth/login"))}render(){return s.Z.isNotEmpty(this.state.token)?(0,r.jsxs)("div",{className:"flex flex-col h-screen bg-slate-700",children:[!this.props.hideNav&&(0,r.jsx)(n.Z,{}),(0,r.jsx)("span",{className:"flex flex-col grow overflow-y-auto ".concat(this.props.className),children:this.props.children})]}):(0,r.jsx)(r.Fragment,{})}constructor(e){super(e),this.state={}}}t.default=(0,o.withRouter)(l)},91580:function(e){e.exports={navbar:"navbar_navbar__2a9Wu",rightSideContainer:"navbar_rightSideContainer__rDPEf",container:"navbar_container__EaLw0",logoContainer:"navbar_logoContainer__cUW1d",logo:"navbar_logo__VRcP0",logoTitle:"navbar_logoTitle__fVGI4",leftSideContainer:"navbar_leftSideContainer__hgHrN",pagesContainer:"navbar_pagesContainer__fbAl_",pageContainer:"navbar_pageContainer__DnRsw",pageLink:"navbar_pageLink__GOJ7T",pageSelected:"navbar_pageSelected__Or1V9",buttonsContainer:"navbar_buttonsContainer__yEhcw"}},2078:function(e,t,a){"use strict";var r=a(67294);let n=r.forwardRef(function({title:e,titleId:t,...a},n){return r.createElement("svg",Object.assign({xmlns:"http://www.w3.org/2000/svg",viewBox:"0 0 20 20",fill:"currentColor","aria-hidden":"true",ref:n,"aria-labelledby":t},a),e?r.createElement("title",{id:t},e):null,r.createElement("path",{d:"M10 8a3 3 0 100-6 3 3 0 000 6zM3.465 14.493a1.23 1.23 0 00.41 1.412A9.957 9.957 0 0010 18c2.31 0 4.438-.784 6.131-2.1.43-.333.604-.903.408-1.41a7.002 7.002 0 00-13.074.003z"}))});t.Z=n}}]);