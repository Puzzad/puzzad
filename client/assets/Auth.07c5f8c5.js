import{S as F,i as W,s as D,H as he,W as N,o as _e,b as u,d as B,f as r,I as ve,J as be,K as we,t as $,v as k,g as C,U as M,V as te,Q as z,M as I,N as T,ab as se,ac as ne,O as P,ad as x,n as q,q as $e,u as ke,h as de,e as p,l as S,p as me,a as g,a9 as Q,r as ze}from"./index.b92c290c.js";import{s as Z}from"./db.737282b8.js";import{t as Y}from"./FlatToast.svelte_svelte_type_style_lang.ded77a3d.js";import{s as ye,l as ge}from"./auth.0042f8b6.js";import{t as ee}from"./title.a27f956a.js";import{g as Be}from"./index.404ba498.js";function pe(a){let e,t;return{c(){e=N("title"),t=M(a[0])},m(n,o){B(n,e,o),r(e,t)},p(n,o){o&1&&te(t,n[0])},d(n){n&&C(e)}}}function Ce(a){let e,t,n,o=a[0]&&pe(a);const s=a[3].default,l=he(s,a,a[2],null);return{c(){e=N("svg"),o&&o.c(),t=_e(),l&&l.c(),u(e,"xmlns","http://www.w3.org/2000/svg"),u(e,"viewBox",a[1]),u(e,"class","svelte-c8tyih")},m(i,c){B(i,e,c),o&&o.m(e,null),r(e,t),l&&l.m(e,null),n=!0},p(i,[c]){i[0]?o?o.p(i,c):(o=pe(i),o.c(),o.m(e,t)):o&&(o.d(1),o=null),l&&l.p&&(!n||c&4)&&ve(l,s,i,i[2],n?we(s,i[2],c,null):be(i[2]),null),(!n||c&2)&&u(e,"viewBox",i[1])},i(i){n||($(l,i),n=!0)},o(i){k(l,i),n=!1},d(i){i&&C(e),o&&o.d(),l&&l.d(i)}}}function Ae(a,e,t){let{$$slots:n={},$$scope:o}=e,{title:s=null}=e,{viewBox:l}=e;return a.$$set=i=>{"title"in i&&t(0,s=i.title),"viewBox"in i&&t(1,l=i.viewBox),"$$scope"in i&&t(2,o=i.$$scope)},[s,l,o,n]}class le extends F{constructor(e){super(),W(this,e,Ae,Ce,D,{title:0,viewBox:1})}}function He(a){let e;return{c(){e=N("path"),u(e,"d","M488 261.8C488 403.3 391.1 504 248 504 110.8 504 0 393.2 0 256S110.8 8 248 8c66.8 0 123 24.5 166.3 64.9l-67.5 64.9C258.5 52.6 94.3 116.6 94.3 256c0 86.5 69.1 156.6 153.7 156.6 98.2 0 135-70.4 140.8-106.9H248v-85.3h236.1c2.3 12.7 3.9 24.9 3.9 41.4z")},m(t,n){B(t,e,n)},p:q,d(t){t&&C(e)}}}function Le(a){let e,t;const n=[{viewBox:"0 0 488 512"},a[0]];let o={$$slots:{default:[He]},$$scope:{ctx:a}};for(let s=0;s<n.length;s+=1)o=z(o,n[s]);return e=new le({props:o}),{c(){I(e.$$.fragment)},m(s,l){T(e,s,l),t=!0},p(s,[l]){const i=l&1?se(n,[n[0],ne(s[0])]):{};l&2&&(i.$$scope={dirty:l,ctx:s}),e.$set(i)},i(s){t||($(e.$$.fragment,s),t=!0)},o(s){k(e.$$.fragment,s),t=!1},d(s){P(e,s)}}}function Se(a,e,t){return a.$$set=n=>{t(0,e=z(z({},e),x(n)))},e=x(e),[e]}class Ve extends F{constructor(e){super(),W(this,e,Se,Le,D,{})}}function Me(a){let e;return{c(){e=N("path"),u(e,"d","M297.216 243.2c0 15.616-11.52 28.416-26.112 28.416-14.336 0-26.112-12.8-26.112-28.416s11.52-28.416 26.112-28.416c14.592 0 26.112 12.8 26.112 28.416zm-119.552-28.416c-14.592 0-26.112 12.8-26.112 28.416s11.776 28.416 26.112 28.416c14.592 0 26.112-12.8 26.112-28.416.256-15.616-11.52-28.416-26.112-28.416zM448 52.736V512c-64.494-56.994-43.868-38.128-118.784-107.776l13.568 47.36H52.48C23.552 451.584 0 428.032 0 398.848V52.736C0 23.552 23.552 0 52.48 0h343.04C424.448 0 448 23.552 448 52.736zm-72.96 242.688c0-82.432-36.864-149.248-36.864-149.248-36.864-27.648-71.936-26.88-71.936-26.88l-3.584 4.096c43.52 13.312 63.744 32.512 63.744 32.512-60.811-33.329-132.244-33.335-191.232-7.424-9.472 4.352-15.104 7.424-15.104 7.424s21.248-20.224 67.328-33.536l-2.56-3.072s-35.072-.768-71.936 26.88c0 0-36.864 66.816-36.864 149.248 0 0 21.504 37.12 78.08 38.912 0 0 9.472-11.52 17.152-21.248-32.512-9.728-44.8-30.208-44.8-30.208 3.766 2.636 9.976 6.053 10.496 6.4 43.21 24.198 104.588 32.126 159.744 8.96 8.96-3.328 18.944-8.192 29.44-15.104 0 0-12.8 20.992-46.336 30.464 7.68 9.728 16.896 20.736 16.896 20.736 56.576-1.792 78.336-38.912 78.336-38.912z")},m(t,n){B(t,e,n)},p:q,d(t){t&&C(e)}}}function Ie(a){let e,t;const n=[{viewBox:"0 0 448 512"},a[0]];let o={$$slots:{default:[Me]},$$scope:{ctx:a}};for(let s=0;s<n.length;s+=1)o=z(o,n[s]);return e=new le({props:o}),{c(){I(e.$$.fragment)},m(s,l){T(e,s,l),t=!0},p(s,[l]){const i=l&1?se(n,[n[0],ne(s[0])]):{};l&2&&(i.$$scope={dirty:l,ctx:s}),e.$set(i)},i(s){t||($(e.$$.fragment,s),t=!0)},o(s){k(e.$$.fragment,s),t=!1},d(s){P(e,s)}}}function Te(a,e,t){return a.$$set=n=>{t(0,e=z(z({},e),x(n)))},e=x(e),[e]}class Pe extends F{constructor(e){super(),W(this,e,Te,Ie,D,{})}}function qe(a){let e;return{c(){e=N("path"),u(e,"d","M40.1 32L10 108.9v314.3h107V480h60.2l56.8-56.8h87l117-117V32H40.1zm357.8 254.1L331 353H224l-56.8 56.8V353H76.9V72.1h321v214zM331 149v116.9h-40.1V149H331zm-107 0v116.9h-40.1V149H224z")},m(t,n){B(t,e,n)},p:q,d(t){t&&C(e)}}}function xe(a){let e,t;const n=[{viewBox:"0 0 448 512"},a[0]];let o={$$slots:{default:[qe]},$$scope:{ctx:a}};for(let s=0;s<n.length;s+=1)o=z(o,n[s]);return e=new le({props:o}),{c(){I(e.$$.fragment)},m(s,l){T(e,s,l),t=!0},p(s,[l]){const i=l&1?se(n,[n[0],ne(s[0])]):{};l&2&&(i.$$scope={dirty:l,ctx:s}),e.$set(i)},i(s){t||($(e.$$.fragment,s),t=!0)},o(s){k(e.$$.fragment,s),t=!1},d(s){P(e,s)}}}function Ee(a,e,t){return a.$$set=n=>{t(0,e=z(z({},e),x(n)))},e=x(e),[e]}class Fe extends F{constructor(e){super(),W(this,e,Ee,xe,D,{})}}function We(a){let e,t,n,o,s,l;return{c(){e=p("p"),t=M("You're already logged in, "),n=p("a"),n.textContent="Logout",o=M("?"),u(n,"href",""),u(e,"class","svelte-1dexrlt")},m(i,c){B(i,e,c),r(e,t),r(e,n),r(e,o),s||(l=S(n,"click",me(ge)),s=!0)},p:q,i:q,o:q,d(i){i&&C(e),s=!1,l()}}}function De(a){let e,t,n,o,s,l,i,c,_,d,O,h,y,U,V,j,A,m,v,ae,G,ie,b,oe,J,re,w,ce,E,R,ue,K,H,X,fe;return i=new Ve({}),d=new Pe({}),y=new Fe({}),{c(){e=p("form"),t=p("div"),n=p("p"),n.textContent="Sign in with one of these third party providers",o=g(),s=p("div"),l=p("div"),I(i.$$.fragment),c=g(),_=p("div"),I(d.$$.fragment),O=g(),h=p("div"),I(y.$$.fragment),U=g(),V=p("p"),V.textContent="or",j=g(),A=p("p"),m=M(a[4]),v=M(" with email"),ae=g(),G=p("label"),G.textContent="Email:",ie=g(),b=p("input"),oe=g(),J=p("label"),J.textContent="Password:",re=g(),w=p("input"),ce=g(),E=p("button"),R=M(a[4]),ue=g(),K=p("section"),u(n,"class","svelte-1dexrlt"),u(l,"class","icon svelte-1dexrlt"),u(_,"class","icon svelte-1dexrlt"),u(h,"class","icon svelte-1dexrlt"),u(s,"class","socialicons svelte-1dexrlt"),u(V,"class","svelte-1dexrlt"),u(A,"class","svelte-1dexrlt"),u(t,"class","socialwrapper svelte-1dexrlt"),u(G,"for","email"),u(b,"id","email"),u(b,"type","email"),u(b,"name","email"),b.required=!0,u(J,"for","password"),u(w,"id","password"),u(w,"type","password"),u(w,"name","password"),w.required=!0,u(E,"type","submit"),E.disabled=a[6],u(e,"class","basic")},m(f,L){B(f,e,L),r(e,t),r(t,n),r(t,o),r(t,s),r(s,l),T(i,l,null),r(s,c),r(s,_),T(d,_,null),r(s,O),r(s,h),T(y,h,null),r(t,U),r(t,V),r(t,j),r(t,A),r(A,m),r(A,v),r(e,ae),r(e,G),r(e,ie),r(e,b),Q(b,a[1]),r(e,oe),r(e,J),r(e,re),r(e,w),Q(w,a[2]),r(e,ce),r(e,E),r(E,R),r(e,ue),r(e,K),K.innerHTML=a[3],H=!0,X||(fe=[S(l,"click",a[9]),S(_,"click",a[10]),S(h,"click",a[11]),S(b,"input",a[12]),S(w,"input",a[13]),S(e,"submit",me(a[0]))],X=!0)},p(f,L){(!H||L&16)&&te(m,f[4]),L&2&&b.value!==f[1]&&Q(b,f[1]),L&4&&w.value!==f[2]&&Q(w,f[2]),(!H||L&16)&&te(R,f[4]),(!H||L&8)&&(K.innerHTML=f[3])},i(f){H||($(i.$$.fragment,f),$(d.$$.fragment,f),$(y.$$.fragment,f),H=!0)},o(f){k(i.$$.fragment,f),k(d.$$.fragment,f),k(y.$$.fragment,f),H=!1},d(f){f&&C(e),P(i),P(d),P(y),X=!1,ze(fe)}}}function Ne(a){let e,t,n,o;const s=[De,We],l=[];function i(c,_){var d;return(d=c[5])!=null&&d.user?1:0}return e=i(a),t=l[e]=s[e](a),{c(){t.c(),n=_e()},m(c,_){l[e].m(c,_),B(c,n,_),o=!0},p(c,[_]){let d=e;e=i(c),e===d?l[e].p(c,_):($e(),k(l[d],1,1,()=>{l[d]=null}),ke(),t=l[e],t?t.p(c,_):(t=l[e]=s[e](c),t.c()),$(t,1),t.m(n.parentNode,n))},i(c){o||($(t),o=!0)},o(c){k(t),o=!1},d(c){l[e].d(c),c&&C(n)}}}function Oe(a,e,t){let n,o;de(a,Be,m=>t(14,n=m)),de(a,ye,m=>t(5,o=m));let{type:s="login"}=e,l,i,c,_,d;const O=async()=>{let m,v;switch(s){case"signup":({error:v}=await Z.auth.signUp({email:l,password:i})),m="An email has been sent to you for verification!";break;case"login":({error:v}=await Z.auth.signInWithPassword({email:l,password:i})),m="Login success, redirecting.",n("/",{});break}v?Y.add({title:"Error",description:v.message,duration:1e4,type:"error"}):Y.add({title:"Success",description:m,duration:1e4,type:"success"})},h=async m=>{const{error:v}=await Z.auth.signInWithOAuth({provider:m});v?Y.add({title:"Error",description:v.message,duration:1e4,type:"error"}):Y.add({title:"Success",description:"Login success, redirecting.",duration:1e4,type:"success"})};switch(s){case"signup":ee.set("Puzzad: Signup"),_="Already have an account? <a href='/login'>Login!</a>",d="Sign up";break;case"login":ee.set("Puzzad: Login"),_="Don't have an account? <a href='/signup'>Sign up!</a>",d="Sign In";break;case"logout":ee.set("Puzzad: Logout"),ge()}const y=()=>h("google"),U=()=>h("discord"),V=()=>h("twitch");function j(){l=this.value,t(1,l)}function A(){i=this.value,t(2,i)}return a.$$set=m=>{"type"in m&&t(8,s=m.type)},[O,l,i,_,d,o,c,h,s,y,U,V,j,A]}class Ye extends F{constructor(e){super(),W(this,e,Oe,Ne,D,{type:8,authAction:0})}get authAction(){return this.$$.ctx[0]}}export{Ye as A};
