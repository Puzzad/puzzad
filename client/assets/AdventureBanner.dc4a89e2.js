import{S as B,i as F,s as M,a3 as O,a4 as S,e as k,M as D,a as p,b as c,j as A,y as I,d as w,f as b,N as V,a5 as Y,t as z,v as E,g as h,O as G,U,V as j}from"./index.b92c290c.js";import{s as H}from"./db.737282b8.js";import{A as J}from"./AdventureLogo.91b8d310.js";function P(l){let e,a;return{c(){e=k("code"),a=U(l[1]),c(e,"class","svelte-1d9cxwq")},m(t,i){w(t,e,i),b(e,a)},p(t,i){i&2&&j(a,t[1])},d(t){t&&h(e)}}}function C(l){let e;return{c(){e=k("div"),e.textContent="ADMIN ONLY",c(e,"class","admin svelte-1d9cxwq")},m(a,t){w(a,e,t)},d(a){a&&h(e)}}}function L(l){let e,a=l[4]===0?"Free!":"\xA3"+l[4],t;return{c(){e=k("div"),t=U(a),c(e,"class","price svelte-1d9cxwq")},m(i,o){w(i,e,o),b(e,t)},p(i,o){o&16&&a!==(a=i[4]===0?"Free!":"\xA3"+i[4])&&j(t,a)},d(i){i&&h(e)}}}function K(l){let e,a,t,i,o,_,m,v,g,s;function y(n){l[6](n)}let q={};l[0]!==void 0&&(q.name=l[0]),t=new J({props:q}),O.push(()=>S(t,"name",y));let u=l[1]&&P(l),f=!l[3]&&C(),r=l[4]!==null&&L(l);return{c(){e=k("a"),a=k("div"),D(t.$$.fragment),o=p(),u&&u.c(),_=p(),f&&f.c(),m=p(),r&&r.c(),c(a,"class","logo svelte-1d9cxwq"),c(e,"class",v=A(l[2].toLowerCase())+" svelte-1d9cxwq"),I(e,"background-image","url('"+l[5]+"')"),c(e,"href",g=l[1]?"/games/"+l[1]:"/adventures/"+l[0])},m(n,d){w(n,e,d),b(e,a),V(t,a,null),b(e,o),u&&u.m(e,null),b(e,_),f&&f.m(e,null),b(e,m),r&&r.m(e,null),s=!0},p(n,[d]){const N={};!i&&d&1&&(i=!0,N.name=n[0],Y(()=>i=!1)),t.$set(N),n[1]?u?u.p(n,d):(u=P(n),u.c(),u.m(e,_)):u&&(u.d(1),u=null),n[3]?f&&(f.d(1),f=null):f||(f=C(),f.c(),f.m(e,m)),n[4]!==null?r?r.p(n,d):(r=L(n),r.c(),r.m(e,null)):r&&(r.d(1),r=null),(!s||d&4&&v!==(v=A(n[2].toLowerCase())+" svelte-1d9cxwq"))&&c(e,"class",v),(!s||d&3&&g!==(g=n[1]?"/games/"+n[1]:"/adventures/"+n[0]))&&c(e,"href",g)},i(n){s||(z(t.$$.fragment,n),s=!0)},o(n){E(t.$$.fragment,n),s=!1},d(n){n&&h(e),G(t),u&&u.d(),f&&f.d(),r&&r.d()}}}function Q(l,e,a){let{adventureName:t}=e,{code:i}=e,{status:o}=e,{isPublic:_=!0}=e,{price:m=null}=e,v=t&&H.storage.from("adventures").getPublicUrl(t+"/background.jpg").data.publicUrl;function g(s){t=s,a(0,t)}return l.$$set=s=>{"adventureName"in s&&a(0,t=s.adventureName),"code"in s&&a(1,i=s.code),"status"in s&&a(2,o=s.status),"isPublic"in s&&a(3,_=s.isPublic),"price"in s&&a(4,m=s.price)},[t,i,o,_,m,v,g]}class X extends B{constructor(e){super(),F(this,e,Q,K,M,{adventureName:0,code:1,status:2,isPublic:3,price:4})}}export{X as A};
