import{S as q,i as B,s as C,h as _,a as d,g as y,t as p,c as w,b as h,d as m,o as L,f as R,j as $,e as j,k as W,l as A,m as M,n as N,p as S,r as Y}from"./index.a596424d.js";import{A as G}from"./AdventureBanner.32afdcc1.js";import{S as H}from"./Spinner.08f32d7c.js";function g(f,a,r){const e=f.slice();return e[4]=a[r],e}function I(f){let a,r,e,l,s=f[0],n=[];for(let t=0;t<s.length;t+=1)n[t]=k(g(f,s,t));const c=t=>p(n[t],1,1,()=>{n[t]=null});let o=null;return s.length||(o=b(f)),{c(){a=_("h2"),a.textContent="Your adventures",r=$();for(let t=0;t<n.length;t+=1)n[t].c();e=j(),o&&o.c()},m(t,i){d(t,a,i),d(t,r,i);for(let u=0;u<n.length;u+=1)n[u].m(t,i);d(t,e,i),o&&o.m(t,i),l=!0},p(t,i){if(i&5){s=t[0];let u;for(u=0;u<s.length;u+=1){const v=g(t,s,u);n[u]?(n[u].p(v,i),h(n[u],1)):(n[u]=k(v),n[u].c(),h(n[u],1),n[u].m(e.parentNode,e))}for(y(),u=s.length;u<n.length;u+=1)c(u);w(),!s.length&&o?o.p(t,i):s.length?o&&(o.d(1),o=null):(o=b(t),o.c(),o.m(e.parentNode,e))}},i(t){if(!l){for(let i=0;i<s.length;i+=1)h(n[i]);l=!0}},o(t){n=n.filter(Boolean);for(let i=0;i<n.length;i+=1)p(n[i]);l=!1},d(t){t&&m(a),t&&m(r),W(n,t),t&&m(e),o&&o.d(t)}}}function J(f){let a,r;return a=new H({}),{c(){A(a.$$.fragment)},m(e,l){M(a,e,l),r=!0},p:N,i(e){r||(h(a.$$.fragment,e),r=!0)},o(e){p(a.$$.fragment,e),r=!1},d(e){S(a,e)}}}function b(f){let a,r,e;return{c(){a=_("blockquote"),a.textContent=`${f[2]}`,r=$(),e=_("p"),e.innerHTML=`You aren&#39;t part of any adventures! You can
                <a href="/#/Adventures">browse the available adventures</a>, or if you
                want to join a friend you can <a href="/#/">enter the game code</a> to
                jump straight in.
            `,Y(a,"class","svelte-1xfkhaw")},m(l,s){d(l,a,s),d(l,r,s),d(l,e,s)},p:N,d(l){l&&m(a),l&&m(r),l&&m(e)}}}function k(f){var e,l;let a,r;return a=new G({props:{price:"",status:f[4].status,adventureName:(l=(e=f[4].adventures)==null?void 0:e.name)!=null?l:"Unknown",code:f[4].code}}),{c(){A(a.$$.fragment)},m(s,n){M(a,s,n),r=!0},p(s,n){var o,t;const c={};n&1&&(c.status=s[4].status),n&1&&(c.adventureName=(t=(o=s[4].adventures)==null?void 0:o.name)!=null?t:"Unknown"),n&1&&(c.code=s[4].code),a.$set(c)},i(s){r||(h(a.$$.fragment,s),r=!0)},o(s){p(a.$$.fragment,s),r=!1},d(s){S(a,s)}}}function K(f){let a,r,e,l;const s=[J,I],n=[];function c(o,t){return o[1]?0:1}return r=c(f),e=n[r]=s[r](f),{c(){a=_("section"),e.c()},m(o,t){d(o,a,t),n[r].m(a,null),l=!0},p(o,[t]){let i=r;r=c(o),r===i?n[r].p(o,t):(y(),p(n[i],1,1,()=>{n[i]=null}),w(),e=n[r],e?e.p(o,t):(e=n[r]=s[r](o),e.c()),h(e,1),e.m(a,null))},i(o){l||(h(e),l=!0)},o(o){p(e),l=!1},d(o){o&&m(a),n[r].d()}}}function T(f,a,r){let e=[],l=!0;L(async function(){let{data:c,error:o}=await R.from("games").select(`
                    id,status, adventures ( name ), status, code
                `);o||(r(1,l=!1),r(0,e=c))});const s=[`\u201CNever say 'no' to adventures. Always say 'yes,' otherwise you'll lead a very dull life.\u201D
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
\u2015 C.S. Lewis`],n=s[Math.floor(Math.random()*s.length)];return[e,l,n]}class E extends q{constructor(a){super(),B(this,a,T,K,C,{})}}export{E as default};
