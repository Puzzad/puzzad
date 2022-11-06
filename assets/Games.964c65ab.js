import{S as R,i as B,s as M,h as _,a as p,g,t as m,c as v,b as d,d as h,o as L,f as j,j as N,e as C,k as T,l as b,m as k,n as S,p as y,r as W}from"./index.7594bded.js";import{A as Y}from"./AdventureBanner.14cafb8e.js";import{S as G}from"./Spinner.a37d58bf.js";import{R as H}from"./RandomText.fb61a7ea.js";function w(f,a,r){const t=f.slice();return t[3]=a[r],t}function I(f){let a,r,t,l,s=f[0],e=[];for(let n=0;n<s.length;n+=1)e[n]=q(w(f,s,n));const c=n=>m(e[n],1,1,()=>{e[n]=null});let o=null;return s.length||(o=A(f)),{c(){a=_("h2"),a.textContent="Your adventures",r=N();for(let n=0;n<e.length;n+=1)e[n].c();t=C(),o&&o.c()},m(n,i){p(n,a,i),p(n,r,i);for(let u=0;u<e.length;u+=1)e[u].m(n,i);p(n,t,i),o&&o.m(n,i),l=!0},p(n,i){if(i&5){s=n[0];let u;for(u=0;u<s.length;u+=1){const $=w(n,s,u);e[u]?(e[u].p($,i),d(e[u],1)):(e[u]=q($),e[u].c(),d(e[u],1),e[u].m(t.parentNode,t))}for(g(),u=s.length;u<e.length;u+=1)c(u);v(),!s.length&&o?o.p(n,i):s.length?o&&(g(),m(o,1,1,()=>{o=null}),v()):(o=A(n),o.c(),d(o,1),o.m(t.parentNode,t))}},i(n){if(!l){for(let i=0;i<s.length;i+=1)d(e[i]);l=!0}},o(n){e=e.filter(Boolean);for(let i=0;i<e.length;i+=1)m(e[i]);l=!1},d(n){n&&h(a),n&&h(r),T(e,n),n&&h(t),o&&o.d(n)}}}function J(f){let a,r;return a=new G({}),{c(){b(a.$$.fragment)},m(t,l){k(a,t,l),r=!0},p:S,i(t){r||(d(a.$$.fragment,t),r=!0)},o(t){m(a.$$.fragment,t),r=!1},d(t){y(a,t)}}}function A(f){let a,r,t,l,s;return r=new H({props:{options:f[2]}}),{c(){a=_("blockquote"),b(r.$$.fragment),t=N(),l=_("p"),l.innerHTML=`You aren&#39;t part of any adventures! You can
                <a href="/#/Adventures">browse the available adventures</a>, or if you
                want to join a friend you can <a href="/#/">enter the game code</a> to
                jump straight in.
            `,W(a,"class","svelte-1xfkhaw")},m(e,c){p(e,a,c),k(r,a,null),p(e,t,c),p(e,l,c),s=!0},p:S,i(e){s||(d(r.$$.fragment,e),s=!0)},o(e){m(r.$$.fragment,e),s=!1},d(e){e&&h(a),y(r),e&&h(t),e&&h(l)}}}function q(f){var t,l;let a,r;return a=new Y({props:{price:"",status:f[3].status,adventureName:(l=(t=f[3].adventures)==null?void 0:t.name)!=null?l:"Unknown",code:f[3].code}}),{c(){b(a.$$.fragment)},m(s,e){k(a,s,e),r=!0},p(s,e){var o,n;const c={};e&1&&(c.status=s[3].status),e&1&&(c.adventureName=(n=(o=s[3].adventures)==null?void 0:o.name)!=null?n:"Unknown"),e&1&&(c.code=s[3].code),a.$set(c)},i(s){r||(d(a.$$.fragment,s),r=!0)},o(s){m(a.$$.fragment,s),r=!1},d(s){y(a,s)}}}function K(f){let a,r,t,l;const s=[J,I],e=[];function c(o,n){return o[1]?0:1}return r=c(f),t=e[r]=s[r](f),{c(){a=_("section"),t.c()},m(o,n){p(o,a,n),e[r].m(a,null),l=!0},p(o,[n]){let i=r;r=c(o),r===i?e[r].p(o,n):(g(),m(e[i],1,1,()=>{e[i]=null}),v(),t=e[r],t?t.p(o,n):(t=e[r]=s[r](o),t.c()),d(t,1),t.m(a,null))},i(o){l||(d(t),l=!0)},o(o){m(t),l=!1},d(o){o&&h(a),e[r].d()}}}function U(f,a,r){let t=[],l=!0;return L(async function(){let{data:e,error:c}=await j.from("games").select(`
                    id,status, adventures ( name ), status, code
                `);c||(r(1,l=!1),r(0,t=e))}),[t,l,[`\u201CNever say 'no' to adventures. Always say 'yes,' otherwise you'll lead a very dull life.\u201D
\u2015 Ian Fleming`,`\u201CIf happiness is the goal \u2013 and it should be, then adventures should be a priority.\u201D
\u2015 Richard Branson`,`\u201CLet us step into the night and pursue that flighty temptress, adventure\u201C
\u2015 J K Rowling`,`\u201CLife is either a daring adventure or nothing at all.\u201D
\u2015 Helen Keller`,`\u201CAdventure is worthwhile in itself.\u201D
\u2015 Amelia Earhart`,`\u201CNever fear quarrels, but seek hazardous adventures.\u201D
\u2015 Alexandre Dumas`,`\u201CWhen you see someone putting on his Big Boots, you can be pretty sure that an Adventure is going to happen.\u201D
\u2015 A.A. Milne`,`\u201CWe are plain quiet folk and have no use for adventures. Nasty disturbing uncomfortable things! Make you late for dinner!\u201D
\u2015 J.R.R. Tolkien`,`\u201CMake your choice, adventurous Stranger,
Strike the bell and bide the danger,
Or wonder, till it drives you mad,
What would have followed if you had.\u201D
\u2015 C.S. Lewis`]]}class O extends R{constructor(a){super(),B(this,a,U,K,M,{})}}export{O as default};