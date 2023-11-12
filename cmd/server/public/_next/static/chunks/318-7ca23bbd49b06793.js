(self.webpackChunk_N_E=self.webpackChunk_N_E||[]).push([[318],{76489:function(t,e){"use strict";/*!
 * cookie
 * Copyright(c) 2012-2014 Roman Shtylman
 * Copyright(c) 2015 Douglas Christopher Wilson
 * MIT Licensed
 */e.Q=function(t,e){if("string"!=typeof t)throw TypeError("argument str must be a string");for(var n={},r=(e||{}).decode||i,a=0;a<t.length;){var o=t.indexOf("=",a);if(-1===o)break;var s=t.indexOf(";",a);if(-1===s)s=t.length;else if(s<o){a=t.lastIndexOf(";",o-1)+1;continue}var c=t.slice(a,o).trim();if(void 0===n[c]){var l=t.slice(o+1,s).trim();34===l.charCodeAt(0)&&(l=l.slice(1,-1)),n[c]=function(t,e){try{return e(t)}catch(e){return t}}(l,r)}a=s+1}return n},e.q=function(t,e,i){var o=i||{},s=o.encode||a;if("function"!=typeof s)throw TypeError("option encode is invalid");if(!r.test(t))throw TypeError("argument name is invalid");var c=s(e);if(c&&!r.test(c))throw TypeError("argument val is invalid");var l=t+"="+c;if(null!=o.maxAge){var u=o.maxAge-0;if(isNaN(u)||!isFinite(u))throw TypeError("option maxAge is invalid");l+="; Max-Age="+Math.floor(u)}if(o.domain){if(!r.test(o.domain))throw TypeError("option domain is invalid");l+="; Domain="+o.domain}if(o.path){if(!r.test(o.path))throw TypeError("option path is invalid");l+="; Path="+o.path}if(o.expires){var d=o.expires;if("[object Date]"!==n.call(d)&&!(d instanceof Date)||isNaN(d.valueOf()))throw TypeError("option expires is invalid");l+="; Expires="+d.toUTCString()}if(o.httpOnly&&(l+="; HttpOnly"),o.secure&&(l+="; Secure"),o.priority)switch("string"==typeof o.priority?o.priority.toLowerCase():o.priority){case"low":l+="; Priority=Low";break;case"medium":l+="; Priority=Medium";break;case"high":l+="; Priority=High";break;default:throw TypeError("option priority is invalid")}if(o.sameSite)switch("string"==typeof o.sameSite?o.sameSite.toLowerCase():o.sameSite){case!0:case"strict":l+="; SameSite=Strict";break;case"lax":l+="; SameSite=Lax";break;case"none":l+="; SameSite=None";break;default:throw TypeError("option sameSite is invalid")}return l};var n=Object.prototype.toString,r=/^[\u0009\u0020-\u007e\u0080-\u00ff]+$/;function i(t){return -1!==t.indexOf("%")?decodeURIComponent(t):t}function a(t){return encodeURIComponent(t)}},96245:function(t,e,n){"use strict";function r(t){this.message=t}r.prototype=Error(),r.prototype.name="InvalidCharacterError";var i="undefined"!=typeof window&&window.atob&&window.atob.bind(window)||function(t){var e=String(t).replace(/=+$/,"");if(e.length%4==1)throw new r("'atob' failed: The string to be decoded is not correctly encoded.");for(var n,i,a=0,o=0,s="";i=e.charAt(o++);~i&&(n=a%4?64*n+i:i,a++%4)&&(s+=String.fromCharCode(255&n>>(-2*a&6))))i="ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/=".indexOf(i);return s};function a(t){this.message=t}a.prototype=Error(),a.prototype.name="InvalidTokenError",e.Z=function(t,e){if("string"!=typeof t)throw new a("Invalid token specified");var n=!0===(e=e||{}).header?0:1;try{return JSON.parse(function(t){var e,n=t.replace(/-/g,"+").replace(/_/g,"/");switch(n.length%4){case 0:break;case 2:n+="==";break;case 3:n+="=";break;default:throw"Illegal base64url string!"}try{return e=n,decodeURIComponent(i(e).replace(/(.)/g,function(t,e){var n=e.charCodeAt(0).toString(16).toUpperCase();return n.length<2&&(n="0"+n),"%"+n}))}catch(t){return i(n)}}(t.split(".")[n]))}catch(t){throw new a("Invalid token specified: "+t.message)}}},81280:function(t,e,n){"use strict";n.d(e,{Z:function(){return r}});class r{static isEmpty(t){return!t||""===t.trim()}static isNotEmpty(t){return!r.isEmpty(t)}static valueOrEmpty(t,e){return r.isEmpty(t)?e:t}static isEmailValid(t){return!!t.trim()&&/^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/.test(t)}static PasswordError(t){let e="";t.length<8&&(e="Password should be at least 8 characters long");let n=/[a-z]+/.test(t);return n||(e="Password should contain at least one lowercase character"),(n=/[A-Z]+/.test(t))||(e="Password should contain at least one uppercase character"),(n=/[0-9]+/.test(t))||(e="Password should contain at least one digit"),(n=/[!@#$%^&*.?-]+/.test(t))||(e="Password should contain at least one special character"),e}}},77546:function(t,e,n){"use strict";n.d(e,{Z:function(){return o}});var r=n(85893);n(67294);var i=n(25675),a=n.n(i);function o(t){var e,n,i,o;let s=null!==(e=t.size)&&void 0!==e?e:90;return t.avatar?(0,r.jsx)(a(),{unoptimized:!0,src:t.avatar,alt:"Avatar",width:s,height:s,style:{objectFit:"cover"},className:"rounded-full aspect-square ".concat(null!==(n=t.className)&&void 0!==n?n:""),onClick:t.onCLick}):(0,r.jsx)("div",{style:{height:s,width:s},className:"aspect-square flex items-center justify-center bg-indigo-100 text-indigo-600 rounded-full text-center ".concat(null!==(i=t.textStyle)&&void 0!==i?i:"text-3xl"," font-bold ring-2 ring-indigo-500 ").concat(null!==(o=t.className)&&void 0!==o?o:""),onClick:t.onCLick,children:t.icon?t.icon:t.avatarTitle})}},31730:function(t,e,n){"use strict";var r=n(67294),i=n(33356);class a extends r.Component{constructor(...t){super(...t),this.cookies=new i.Z,this.logout=t=>{this.cookies.remove("token"),void 0!==t&&t()}}}e.Z=a},43067:function(t,e,n){"use strict";n.d(e,{Z:function(){return o}});var r=n(85893),i=n(83365),a=n.n(i);function o(t){var e;let n="full"===t.width?a().btnFull:"",i=e=>{e.preventDefault(),t.onClick&&t.onClick()};return(0,r.jsx)("button",{type:t.type,className:"".concat(a().btn," ").concat(n," ").concat((()=>{if("primary"===t.variant)return t.disabled?a().btnPrimaryDisabled:a().btnPrimary;if("secondary"===t.variant)return t.disabled?a().btnSecondaryDisabled:a().btnSecondary;if("group"===t.variant){let e=t.selected?a().btnPrimary:a().btnSecondary;return"".concat(a().btnGroup," ").concat(e)}return""})()," ").concat(null!==(e=t.className)&&void 0!==e?e:""),onClick:i,disabled:t.disabled,children:t.title},t.title)}},84639:function(t,e,n){"use strict";n.d(e,{Z:function(){return a}});var r=n(85893);n(67294);var i=n(81280);function a(t){return i.Z.isNotEmpty(t.message)?(0,r.jsxs)("div",{className:"p-4 mb-4 text-sm text-red-800 rounded-lg bg-red-50 dark:bg-slate-700 dark:text-red-400 ".concat(t.className?t.className:""),role:"alert",children:[(0,r.jsx)("span",{className:"font-medium",children:"Error :"})," ",t.message]}):(0,r.jsx)(r.Fragment,{})}},27888:function(t,e,n){"use strict";n.d(e,{Z:function(){return v}});var r=n(85893),i=n(67294),a=n(41664),o=n.n(a),s=n(25675),c=n.n(s);class l{constructor(t,e){this.href=t,this.title=e}}var u=n(91580),d=n.n(u),h=n(77546),p=n(2078),f=n(11163);let g=i.forwardRef(function({title:t,titleId:e,...n},r){return i.createElement("svg",Object.assign({xmlns:"http://www.w3.org/2000/svg",fill:"none",viewBox:"0 0 24 24",strokeWidth:1.5,stroke:"currentColor","aria-hidden":"true",ref:r,"aria-labelledby":e},n),t?i.createElement("title",{id:e},t):null,i.createElement("path",{strokeLinecap:"round",strokeLinejoin:"round",d:"M15.75 9V5.25A2.25 2.25 0 0013.5 3h-6a2.25 2.25 0 00-2.25 2.25v13.5A2.25 2.25 0 007.5 21h6a2.25 2.25 0 002.25-2.25V15M12 9l-3 3m0 0l3 3m-3-3h12.75"}))});var m=n(31730);class b extends m.Z{get pages(){let t=[];return this.props.hideAllPages||(t.push(new l("/","Pets")),t.push(new l("/agenda","Agenda")),this.props.pages&&t.push(...this.props.pages)),t}render(){return(0,r.jsx)("nav",{className:d().navbar,children:(0,r.jsxs)("div",{className:d().container,children:[(0,r.jsx)("div",{className:d().rightSideContainer,children:(0,r.jsxs)(o(),{href:"/",className:d().logoContainer,children:[(0,r.jsx)(c(),{src:"/logo.png",width:33,height:35,className:d().logo,alt:"Pet Journal Logo"}),(0,r.jsx)("span",{className:d().logoTitle,children:"Pet Journal"})]})}),(0,r.jsxs)("div",{className:d().leftSideContainer,children:[(0,r.jsx)("div",{className:d().pagesContainer,children:this.pages.map(t=>(0,r.jsx)(o(),{href:t.href,className:d().pageContainer,children:(0,r.jsx)("span",{className:"".concat(d().pageLink," ").concat(this.props.router.pathname===t.href?d().pageSelected:""),children:t.title})},t.title))}),!this.props.hideAllPages&&(0,r.jsxs)("div",{className:"flex flex-row gap-2 items-center",children:[(0,r.jsx)(h.Z,{size:35,icon:(0,r.jsx)(p.Z,{className:"flex p-1"}),onCLick:()=>this.props.router.push("/account"),className:"hover:bg-indigo-200 hover:text-indigo-700"}),(0,r.jsx)(g,{className:"h-10 w-10 text-slate-300 p-2 rounded-full hover:bg-gray-600",onClick:()=>this.logout(()=>this.props.router.replace("/auth/login"))})]}),this.props.buttons&&(0,r.jsx)("div",{className:d().buttonsContainer,children:this.props.buttons})]})]})})}constructor(t){super(t),this.state={}}}var v=(0,f.withRouter)(b)},95546:function(t,e,n){"use strict";n.d(e,{Z:function(){return o}});var r=n(85893);n(67294);var i=n(83815),a=n.n(i);function o(t){var e;let n="full"===t.width?a().textInputFull:"",i=e=>{let n=e.currentTarget.value;t.onInput(n)};return(0,r.jsxs)("div",{className:a().textInputContainer,children:[(0,r.jsx)("label",{htmlFor:t.id,className:t.showLabel?a().label:"sr-only",children:t.placeholder}),(0,r.jsx)("input",{required:t.required,id:t.id,name:t.name,type:t.type,max:t.max,min:t.min,autoComplete:t.autoComplete,className:"".concat(a().textInput," ").concat(n," ").concat(null!==(e=t.classNames)&&void 0!==e?e:""," ").concat(t.hasError&&a().errorHighlight),placeholder:t.placeholder,value:t.value,onInput:i,disabled:t.disabled,onBlur:t.onBlur,autoFocus:t.autoFocus,style:{colorScheme:"dark"}}),(0,r.jsx)("span",{className:a().errorMessage,children:t.hasError?t.errorMessage:""})]})}},43432:function(t,e,n){"use strict";function r(t){return fetch("".concat("http://localhost:8080/api","/auth/login"),{method:"POST",headers:{"Content-Type":"application/json"},body:JSON.stringify(t)})}function i(t){return fetch("".concat("http://localhost:8080/api","/auth/register"),{method:"POST",headers:{"Content-Type":"application/json"},body:JSON.stringify(t.fields)})}function a(t,e){return fetch("".concat("http://localhost:8080/api","/user"),{method:"PATCH",headers:{"Content-Type":"application/json",Authorization:"Bearer ".concat(e)},body:JSON.stringify(t.fields)})}function o(t){return fetch("".concat("http://localhost:8080/api","/user"),{method:"GET",headers:{Authorization:"Bearer ".concat(t)}})}function s(t){return fetch("".concat("http://localhost:8080/api","/user"),{method:"DELETE",headers:{Authorization:"Bearer ".concat(t)}})}function c(){return fetch("".concat("http://localhost:8080/api","/vets"),{method:"GET",headers:{"Content-Type":"application/json"}})}n.d(e,{HN:function(){return r},iD:function(){return a},l5:function(){return s},lx:function(){return i},ui:function(){return c},wB:function(){return o}})},83365:function(t){t.exports={btn:"button_btn__w__d1",btnFull:"button_btnFull__Q8iU8",btnPrimary:"button_btnPrimary__b_fGY",btnSecondary:"button_btnSecondary__Vaxi0",btnGroup:"button_btnGroup__zeRjm",btnPrimaryDisabled:"button_btnPrimaryDisabled__9IQhk",btnSecondaryDisabled:"button_btnSecondaryDisabled__eLBdY"}},91580:function(t){t.exports={navbar:"navbar_navbar__2a9Wu",rightSideContainer:"navbar_rightSideContainer__rDPEf",container:"navbar_container__EaLw0",logoContainer:"navbar_logoContainer__cUW1d",logo:"navbar_logo__VRcP0",logoTitle:"navbar_logoTitle__fVGI4",leftSideContainer:"navbar_leftSideContainer__hgHrN",pagesContainer:"navbar_pagesContainer__fbAl_",pageContainer:"navbar_pageContainer__DnRsw",pageLink:"navbar_pageLink__GOJ7T",pageSelected:"navbar_pageSelected__Or1V9",buttonsContainer:"navbar_buttonsContainer__yEhcw"}},83815:function(t){t.exports={textInput:"textInput_textInput__rwYXO",textInputText:"textInput_textInputText__RbFSM",label:"textInput_label__IpSov",textInputFull:"textInput_textInputFull__BtdFB",errorHighlight:"textInput_errorHighlight__6WM_d",errorMessage:"textInput_errorMessage__2Uw_U",textInputContainer:"textInput_textInputContainer__hcS5_",phoneInput:"textInput_phoneInput__ILwRV",phoneInputFull:"textInput_phoneInputFull__VL9NH"}},11163:function(t,e,n){t.exports=n(96885)},33356:function(t,e,n){"use strict";n.d(e,{Z:function(){return a}});var r=n(76489);function i(t,e={}){let n=t&&"j"===t[0]&&":"===t[1]?t.substr(2):t;if(!e.doNotParse)try{return JSON.parse(n)}catch(t){}return t}var a=class{constructor(t,e={}){var n;this.changeListeners=[],this.HAS_DOCUMENT_COOKIE=!1,this.update=()=>{if(!this.HAS_DOCUMENT_COOKIE)return;let t=this.cookies;this.cookies=r.Q(document.cookie),this._checkChanges(t)};let i="undefined"==typeof document?"":document.cookie;this.cookies="string"==typeof(n=t||i)?r.Q(n):"object"==typeof n&&null!==n?n:{},this.defaultSetOptions=e,this.HAS_DOCUMENT_COOKIE="object"==typeof document&&"string"==typeof document.cookie}_emitChange(t){for(let e=0;e<this.changeListeners.length;++e)this.changeListeners[e](t)}_checkChanges(t){let e=new Set(Object.keys(t).concat(Object.keys(this.cookies)));e.forEach(e=>{t[e]!==this.cookies[e]&&this._emitChange({name:e,value:i(t[e])})})}_startPolling(){this.pollingInterval=setInterval(this.update,300)}_stopPolling(){this.pollingInterval&&clearInterval(this.pollingInterval)}get(t,e={}){return e.doNotUpdate||this.update(),i(this.cookies[t],e)}getAll(t={}){t.doNotUpdate||this.update();let e={};for(let n in this.cookies)e[n]=i(this.cookies[n],t);return e}set(t,e,n){n=n?Object.assign(Object.assign({},this.defaultSetOptions),n):this.defaultSetOptions;let i="string"==typeof e?e:JSON.stringify(e);this.cookies=Object.assign(Object.assign({},this.cookies),{[t]:i}),this.HAS_DOCUMENT_COOKIE&&(document.cookie=r.q(t,i,n)),this._emitChange({name:t,value:e,options:n})}remove(t,e){let n=e=Object.assign(Object.assign({},e),{expires:new Date(1970,1,1,0,0,1),maxAge:0});this.cookies=Object.assign({},this.cookies),delete this.cookies[t],this.HAS_DOCUMENT_COOKIE&&(document.cookie=r.q(t,"",n)),this._emitChange({name:t,value:void 0,options:e})}addChangeListener(t){this.changeListeners.push(t),1===this.changeListeners.length&&("object"==typeof window&&"cookieStore"in window?window.cookieStore.addEventListener("change",this.update):this._startPolling())}removeChangeListener(t){let e=this.changeListeners.indexOf(t);e>=0&&this.changeListeners.splice(e,1),0===this.changeListeners.length&&("object"==typeof window&&"cookieStore"in window?window.cookieStore.removeEventListener("change",this.update):this._stopPolling())}}},2078:function(t,e,n){"use strict";var r=n(67294);let i=r.forwardRef(function({title:t,titleId:e,...n},i){return r.createElement("svg",Object.assign({xmlns:"http://www.w3.org/2000/svg",viewBox:"0 0 20 20",fill:"currentColor","aria-hidden":"true",ref:i,"aria-labelledby":e},n),t?r.createElement("title",{id:e},t):null,r.createElement("path",{d:"M10 8a3 3 0 100-6 3 3 0 000 6zM3.465 14.493a1.23 1.23 0 00.41 1.412A9.957 9.957 0 0010 18c2.31 0 4.438-.784 6.131-2.1.43-.333.604-.903.408-1.41a7.002 7.002 0 00-13.074.003z"}))});e.Z=i}}]);