import{S as M,i as O,s as P,h as c,j as b,x as z,l as U,r as i,a as T,w as l,L as w,m as j,C,M as B,b as k,t as A,d as F,p as E,O as D,f as G}from"./index.7594bded.js";import{T as H,F as I,t as L}from"./FlatToast.2ea36fa8.js";function J(r){let t,n;return t=new I({props:{data:r[4]}}),{c(){U(t.$$.fragment)},m(s,a){j(t,s,a),n=!0},p(s,a){const u={};a&16&&(u.data=s[4]),t.$set(u)},i(s){n||(k(t.$$.fragment,s),n=!0)},o(s){A(t.$$.fragment,s),n=!1},d(s){E(t,s)}}}function K(r){let t,n,s,a,u,d,g,o,h,f,_,S,m,$,v,y;return m=new H({props:{placement:"bottom-right",theme:"dark",$$slots:{default:[J,({data:e})=>({4:e}),({data:e})=>e?16:0]},$$scope:{ctx:r}}}),{c(){t=c("form"),n=c("label"),n.textContent="Email:",s=b(),a=c("input"),u=b(),d=c("label"),d.textContent="Password:",g=b(),o=c("input"),h=b(),f=c("button"),_=z("Sign Up"),S=b(),U(m.$$.fragment),i(n,"for","email"),i(a,"id","email"),i(a,"type","email"),i(a,"name","email"),a.required=!0,i(d,"for","password"),i(o,"id","password"),i(o,"type","password"),i(o,"name","password"),o.required=!0,i(f,"type","submit"),f.disabled=r[2],f.value="Sign up",i(t,"class","basic")},m(e,p){T(e,t,p),l(t,n),l(t,s),l(t,a),w(a,r[0]),l(t,u),l(t,d),l(t,g),l(t,o),w(o,r[1]),l(t,h),l(t,f),l(f,_),T(e,S,p),j(m,e,p),$=!0,v||(y=[C(a,"input",r[5]),C(o,"input",r[6]),C(t,"submit",B(r[3]))],v=!0)},p(e,[p]){p&1&&a.value!==e[0]&&w(a,e[0]),p&2&&o.value!==e[1]&&w(o,e[1]),(!$||p&4)&&(f.disabled=e[2]);const q={};p&144&&(q.$$scope={dirty:p,ctx:e}),m.$set(q)},i(e){$||(k(m.$$.fragment,e),$=!0)},o(e){A(m.$$.fragment,e),$=!1},d(e){e&&F(t),e&&F(S),E(m,e),v=!1,D(y)}}}function N(r,t,n){let s="",a="",u=!1,d={};const g=async()=>{n(2,u=!0);const{_:f,error:_}=await G.auth.signUp({email:s,password:a});n(2,u=!1),_?L.add({title:"Signup",description:_.message,duration:1e4,type:"error"}):L.add({title:"Signup",description:"An email has been sent to you for verification!",duration:1e4,type:"success"})};function o(){s=this.value,n(0,s)}function h(){a=this.value,n(1,a)}return[s,a,u,g,d,o,h]}class V extends M{constructor(t){super(),O(this,t,N,K,P,{})}}export{V as default};