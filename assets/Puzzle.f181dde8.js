import{S as fe,i as ce,s as de,e as Se,a as h,n as V,d as p,o as Te,h as v,u as I,H as K,k as Me,w as C,l as j,j as T,x as q,m as B,I as Ne,J as Pe,b as D,t as H,p as L,z as Q,v as ge,A as ie,G as re,F as $,g as qe,c as Ge,K as De,L as le,q as ue,r as he,M as pe,N as Ae,y as be,O as He,P as Oe}from"./index.fd81a935.js";import{R as Re}from"./RandomText.26b73924.js";import{S as Ue}from"./Spinner.b17cd6d4.js";import{t as Fe,T as Ee,F as We}from"./FlatToast.be15d332.js";function ze(l,e,t){const n=l.slice();return n[18]=e[t],n}function ke(l){let e,t={length:l[6]},n=[];for(let i=0;i<t.length;i+=1)n[i]=ye(ze(l,t,i));return{c(){e=v("div");for(let i=0;i<n.length;i+=1)n[i].c();I(e,"class","confetti-holder svelte-io58ff"),K(e,"rounded",l[9]),K(e,"cone",l[10]),K(e,"no-gravity",l[11])},m(i,a){h(i,e,a);for(let s=0;s<n.length;s+=1)n[s].m(e,null)},p(i,a){if(a&20991){t={length:i[6]};let s;for(s=0;s<t.length;s+=1){const u=ze(i,t,s);n[s]?n[s].p(u,a):(n[s]=ye(u),n[s].c(),n[s].m(e,null))}for(;s<n.length;s+=1)n[s].d(1);n.length=t.length}a&512&&K(e,"rounded",i[9]),a&1024&&K(e,"cone",i[10]),a&2048&&K(e,"no-gravity",i[11])},d(i){i&&p(e),Me(n,i)}}}function ye(l){let e;return{c(){e=v("div"),I(e,"class","confetti svelte-io58ff"),C(e,"--fall-distance",l[8]),C(e,"--size",l[0]+"px"),C(e,"--color",l[14]()),C(e,"--skew",M(-45,45)+"deg,"+M(-45,45)+"deg"),C(e,"--rotation-xyz",M(-10,10)+", "+M(-10,10)+", "+M(-10,10)),C(e,"--rotation-deg",M(0,360)+"deg"),C(e,"--translate-y-multiplier",M(l[2][0],l[2][1])),C(e,"--translate-x-multiplier",M(l[1][0],l[1][1])),C(e,"--scale",.1*M(2,10)),C(e,"--transition-duration",l[4]?`calc(${l[3]}ms * var(--scale))`:`${l[3]}ms`),C(e,"--transition-delay",M(l[5][0],l[5][1])+"ms"),C(e,"--transition-iteration-count",l[4]?"infinite":l[7]),C(e,"--x-spread",1-l[12])},m(t,n){h(t,e,n)},p(t,n){n&256&&C(e,"--fall-distance",t[8]),n&1&&C(e,"--size",t[0]+"px"),n&4&&C(e,"--translate-y-multiplier",M(t[2][0],t[2][1])),n&2&&C(e,"--translate-x-multiplier",M(t[1][0],t[1][1])),n&24&&C(e,"--transition-duration",t[4]?`calc(${t[3]}ms * var(--scale))`:`${t[3]}ms`),n&32&&C(e,"--transition-delay",M(t[5][0],t[5][1])+"ms"),n&144&&C(e,"--transition-iteration-count",t[4]?"infinite":t[7]),n&4096&&C(e,"--x-spread",1-t[12])},d(t){t&&p(e)}}}function Ye(l){let e,t=!l[13]&&ke(l);return{c(){t&&t.c(),e=Se()},m(n,i){t&&t.m(n,i),h(n,e,i)},p(n,[i]){n[13]?t&&(t.d(1),t=null):t?t.p(n,i):(t=ke(n),t.c(),t.m(e.parentNode,e))},i:V,o:V,d(n){t&&t.d(n),n&&p(e)}}}function M(l,e){return Math.random()*(e-l)+l}function je(l,e,t){let{size:n=10}=e,{x:i=[-.5,.5]}=e,{y:a=[.25,1]}=e,{duration:s=2e3}=e,{infinite:u=!1}=e,{delay:f=[0,50]}=e,{colorRange:_=[0,360]}=e,{colorArray:r=[]}=e,{amount:m=50}=e,{iterationCount:b=1}=e,{fallDistance:k="100px"}=e,{rounded:y=!1}=e,{cone:S=!1}=e,{noGravity:z=!1}=e,{xSpread:W=.15}=e,{destroyOnComplete:N=!0}=e,F=!1;Te(()=>{!N||u||b=="infinite"||setTimeout(()=>t(13,F=!0),(s+f[1])*b)});function R(){return r.length?r[Math.round(Math.random()*(r.length-1))]:`hsl(${Math.round(M(_[0],_[1]))}, 75%, 50%`}return l.$$set=c=>{"size"in c&&t(0,n=c.size),"x"in c&&t(1,i=c.x),"y"in c&&t(2,a=c.y),"duration"in c&&t(3,s=c.duration),"infinite"in c&&t(4,u=c.infinite),"delay"in c&&t(5,f=c.delay),"colorRange"in c&&t(15,_=c.colorRange),"colorArray"in c&&t(16,r=c.colorArray),"amount"in c&&t(6,m=c.amount),"iterationCount"in c&&t(7,b=c.iterationCount),"fallDistance"in c&&t(8,k=c.fallDistance),"rounded"in c&&t(9,y=c.rounded),"cone"in c&&t(10,S=c.cone),"noGravity"in c&&t(11,z=c.noGravity),"xSpread"in c&&t(12,W=c.xSpread),"destroyOnComplete"in c&&t(17,N=c.destroyOnComplete)},[n,i,a,s,u,f,m,b,k,y,S,z,W,F,R,_,r,N]}class Be extends fe{constructor(e){super(),ce(this,e,je,Ye,de,{size:0,x:1,y:2,duration:3,infinite:4,delay:5,colorRange:15,colorArray:16,amount:6,iterationCount:7,fallDistance:8,rounded:9,cone:10,noGravity:11,xSpread:12,destroyOnComplete:17})}}function ve(l,e,t){const n=l.slice();return n[7]=e[t],n}function Le(l){let e=l[7].message+"",t;return{c(){t=Q(e)},m(n,i){h(n,t,i)},p(n,i){i&1&&e!==(e=n[7].message+"")&&ie(t,e)},d(n){n&&p(t)}}}function Je(l){let e;return{c(){e=Q(`Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut
                    labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco
                    laboris nisi ut aliquip ex ea commodo consequat.`)},m(t,n){h(t,e,n)},p:V,d(t){t&&p(e)}}}function Ce(l){let e,t,n;function i(){return l[6](l[7])}return{c(){e=v("button"),e.textContent="Reveal this hint",I(e,"class","svelte-6e5e0o")},m(a,s){h(a,e,s),t||(n=$(e,"click",i),t=!0)},p(a,s){l=a},d(a){a&&p(e),t=!1,n()}}}function we(l,e){let t,n=e[7].title+"",i,a,s,u,f,_,r;function m(S,z){return S[7].locked?Je:Le}let b=m(e),k=b(e),y=e[7].locked&&Ce(e);return{key:l,first:null,c(){t=v("h4"),i=Q(n),a=T(),s=v("div"),u=v("p"),k.c(),_=T(),y&&y.c(),r=T(),I(t,"class","svelte-6e5e0o"),I(u,"class",f=ge(e[7].locked?"locked":"")+" svelte-6e5e0o"),I(s,"class","svelte-6e5e0o"),this.first=t},m(S,z){h(S,t,z),q(t,i),h(S,a,z),h(S,s,z),q(s,u),k.m(u,null),q(s,_),y&&y.m(s,null),q(s,r)},p(S,z){e=S,z&1&&n!==(n=e[7].title+"")&&ie(i,n),b===(b=m(e))&&k?k.p(e,z):(k.d(1),k=b(e),k&&(k.c(),k.m(u,null))),z&1&&f!==(f=ge(e[7].locked?"locked":"")+" svelte-6e5e0o")&&I(u,"class",f),e[7].locked?y?y.p(e,z):(y=Ce(e),y.c(),y.m(s,r)):y&&(y.d(1),y=null)},d(S){S&&p(t),S&&p(a),S&&p(s),k.d(),y&&y.d()}}}function Ke(l){let e,t,n,i,a=[],s=new Map,u;n=new Re({props:{options:l[2]}});let f=l[0];const _=r=>r[7].id;for(let r=0;r<f.length;r+=1){let m=ve(l,f,r),b=_(m);s.set(b,a[r]=we(b,m))}return{c(){e=v("section"),t=v("p"),j(n.$$.fragment),i=T();for(let r=0;r<a.length;r+=1)a[r].c();I(e,"class","svelte-6e5e0o")},m(r,m){h(r,e,m),q(e,t),B(n,t,null),q(e,i);for(let b=0;b<a.length;b+=1)a[b].m(e,null);u=!0},p(r,[m]){m&3&&(f=r[0],a=Ne(a,m,_,1,r,f,s,e,Pe,we,null,ve))},i(r){u||(D(n.$$.fragment,r),u=!0)},o(r){H(n.$$.fragment,r),u=!1},d(r){r&&p(e),L(n);for(let m=0;m<a.length;m+=1)a[m].d()}}}function Qe(l,e,t){let{gameCode:n=""}=e,{puzzleId:i=0}=e;const a=async()=>{const r=await re(n),{data:m,error:b}=await r.rpc("gethints",{puzzleid:i,gamecode:n});if(b)throw b;t(0,u=m)},s=async function(r){await(await re(n)).rpc("requesthint",{puzzleid:i,gamecode:n,hintid:r})};let u=[];const f=["Need a hint? Browse our extensive collection below!","Not sure where to go? Try one of our finest hints!","Get your hints here! Freshly plucked from the hint tree!","Psst... Can I interest you in a hint?","We've got the finest hints in all the land. Don't believe me? Try one for free!","Struggling? In a rush? Why not take one of our hand-picked, artisanal hints?"],_=r=>s(r.id);return l.$$set=r=>{"gameCode"in r&&t(3,n=r.gameCode),"puzzleId"in r&&t(4,i=r.puzzleId)},l.$$.update=()=>{l.$$.dirty&24&&n&&i&&a()},[u,s,f,n,i,a,_]}class Ve extends fe{constructor(e){super(),ce(this,e,Qe,Ke,de,{gameCode:3,puzzleId:4,refresh:5})}get refresh(){return this.$$.ctx[5]}}function Xe(l){let e,t=l[6].adventure.name+"",n,i,a=l[6].title+"",s,u,f,_=l[6].content+"",r,m,b,k,y,S,z,W,N,F,R,c,X,Z,P,J,E,ee,x,Y,d,w,O,A;function se(o){l[13](o)}function te(o){l[14](o)}let oe={};l[0].code!==void 0&&(oe.gameCode=l[0].code),l[0].puzzle!==void 0&&(oe.puzzleId=l[0].puzzle),R=new Ve({props:oe}),ue.push(()=>he(R,"gameCode",se)),ue.push(()=>he(R,"puzzleId",te)),l[15](R),E=new Re({props:{options:l[10]}});function me(o,g){return o[6].next?$e:xe}let ne=me(l),U=ne(l),G=l[2]&&Ie();return d=new Ee({props:{placement:"bottom-right",theme:"dark",$$slots:{default:[et,({data:o})=>({6:o}),({data:o})=>o?64:0]},$$scope:{ctx:l}}}),{c(){e=v("h2"),n=Q(t),i=Q(": "),s=Q(a),u=T(),f=new Oe(!1),r=T(),m=v("section"),b=v("form"),k=v("fieldset"),y=v("legend"),y.textContent="Enter a guess",S=T(),z=v("input"),W=T(),N=v("input"),F=T(),j(R.$$.fragment),Z=T(),P=v("dialog"),J=v("h3"),j(E.$$.fragment),ee=T(),U.c(),x=T(),G&&G.c(),Y=T(),j(d.$$.fragment),f.a=r,I(z,"type","text"),z.disabled=l[3],I(z,"class","svelte-ft8dai"),I(N,"type","submit"),N.value="Submit",N.disabled=l[3],I(k,"class","svelte-ft8dai"),I(m,"class","answer svelte-ft8dai"),I(J,"class","svelte-ft8dai"),P.open=l[2],I(P,"class","svelte-ft8dai")},m(o,g){h(o,e,g),q(e,n),q(e,i),q(e,s),h(o,u,g),f.m(_,o,g),h(o,r,g),h(o,m,g),q(m,b),q(b,k),q(k,y),q(k,S),q(k,z),pe(z,l[1]),q(k,W),q(k,N),h(o,F,g),B(R,o,g),h(o,Z,g),h(o,P,g),q(P,J),B(E,J,null),q(P,ee),U.m(P,null),h(o,x,g),G&&G.m(o,g),h(o,Y,g),B(d,o,g),w=!0,O||(A=[$(z,"input",l[11]),$(b,"submit",Ae(l[12]))],O=!0)},p(o,g){(!w||g&64)&&t!==(t=o[6].adventure.name+"")&&ie(n,t),(!w||g&64)&&a!==(a=o[6].title+"")&&ie(s,a),(!w||g&64)&&_!==(_=o[6].content+"")&&f.p(_),(!w||g&8)&&(z.disabled=o[3]),g&2&&z.value!==o[1]&&pe(z,o[1]),(!w||g&8)&&(N.disabled=o[3]);const ae={};!c&&g&1&&(c=!0,ae.gameCode=o[0].code,be(()=>c=!1)),!X&&g&1&&(X=!0,ae.puzzleId=o[0].puzzle,be(()=>X=!1)),R.$set(ae),ne===(ne=me(o))&&U?U.p(o,g):(U.d(1),U=ne(o),U&&(U.c(),U.m(P,null))),(!w||g&4)&&(P.open=o[2]),o[2]?G?g&4&&D(G,1):(G=Ie(),G.c(),D(G,1),G.m(Y.parentNode,Y)):G&&(qe(),H(G,1,1,()=>{G=null}),Ge());const _e={};g&134217792&&(_e.$$scope={dirty:g,ctx:o}),d.$set(_e)},i(o){w||(D(R.$$.fragment,o),D(E.$$.fragment,o),D(G),D(d.$$.fragment,o),w=!0)},o(o){H(R.$$.fragment,o),H(E.$$.fragment,o),H(G),H(d.$$.fragment,o),w=!1},d(o){o&&p(e),o&&p(u),o&&f.d(),o&&p(r),o&&p(m),o&&p(F),l[15](null),L(R,o),o&&p(Z),o&&p(P),L(E),U.d(),o&&p(x),G&&G.d(o),o&&p(Y),L(d,o),O=!1,He(A)}}}function Ze(l){let e,t;return e=new Ue({}),{c(){j(e.$$.fragment)},m(n,i){B(e,n,i),t=!0},p:V,i(n){t||(D(e.$$.fragment,n),t=!0)},o(n){H(e.$$.fragment,n),t=!1},d(n){L(e,n)}}}function xe(l){let e,t,n,i,a;return{c(){e=v("p"),e.textContent="You have completed the adventure!",t=T(),n=v("button"),n.textContent="Continue \xBB",I(e,"class","svelte-ft8dai"),I(n,"class","svelte-ft8dai")},m(s,u){h(s,e,u),h(s,t,u),h(s,n,u),i||(a=$(n,"click",l[17]),i=!0)},p:V,d(s){s&&p(e),s&&p(t),s&&p(n),i=!1,a()}}}function $e(l){let e,t,n,i,a;return{c(){e=v("p"),e.textContent="You have completed this step in the adventure!",t=T(),n=v("button"),n.textContent="Next puzzle \xBB",I(e,"class","svelte-ft8dai"),I(n,"class","svelte-ft8dai")},m(s,u){h(s,e,u),h(s,t,u),h(s,n,u),i||(a=$(n,"click",l[16]),i=!0)},p:V,d(s){s&&p(e),s&&p(t),s&&p(n),i=!1,a()}}}function Ie(l){let e,t,n;return t=new Be({props:{x:[-5,5],y:[0,.1],delay:[0,2e3],infinite:!0,duration:"5000",amount:"400",fallDistance:"110vh"}}),{c(){e=v("div"),j(t.$$.fragment),I(e,"class","confetti-bg svelte-ft8dai")},m(i,a){h(i,e,a),B(t,e,null),n=!0},i(i){n||(D(t.$$.fragment,i),n=!0)},o(i){H(t.$$.fragment,i),n=!1},d(i){i&&p(e),L(t)}}}function et(l){let e,t;return e=new We({props:{data:l[6]}}),{c(){j(e.$$.fragment)},m(n,i){B(e,n,i),t=!0},p(n,i){const a={};i&64&&(a.data=n[6]),e.$set(a)},i(n){t||(D(e.$$.fragment,n),t=!0)},o(n){H(e.$$.fragment,n),t=!1},d(n){L(e,n)}}}function tt(l){let e,t,n,i;const a=[Ze,Xe],s=[];function u(f,_){return f[4]?0:1}return e=u(l),t=s[e]=a[e](l),{c(){t.c(),n=Se()},m(f,_){s[e].m(f,_),h(f,n,_),i=!0},p(f,[_]){let r=e;e=u(f),e===r?s[e].p(f,_):(qe(),H(s[r],1,1,()=>{s[r]=null}),Ge(),t=s[e],t?t.p(f,_):(t=s[e]=a[e](f),t.c()),D(t,1),t.m(n.parentNode,n))},i(f){i||(D(t),i=!0)},o(f){H(t),i=!1},d(f){s[e].d(f),f&&p(n)}}}function nt(l,e,t){let{params:n={}}=e,i={},a="",s=!1,u=!1,f=!0,_=null,r=null,m;const b=()=>{try{k(),re(n.code).then(d=>(r=d,d.from("puzzles").select("title, content, next, storage_slug, adventure (name)").eq("id",n.puzzle))).then(y).then(S).then(z).then(d=>t(6,i=d)).then(()=>{_||W(),t(4,f=!1)}).catch(()=>le("/game/"+n.code))}catch{le("/game/"+n.code)}},k=()=>{console.log("Resetting..."),t(4,f=!0),t(2,s=!1),t(1,a=""),t(3,u=!1)},y=({data:d,error:w})=>{if(w)throw w;if(d.length===0)throw"Puzzle not found";return d[0]},S=d=>{let w=[d];const O=d.content.match(/\$[^$]+?\$/g);return O&&O.forEach(A=>{w.push(A,r.storage.from("puzzles").createSignedUrl(d.storage_slug+"/"+A.substring(1,A.length-1),60*60))}),Promise.all(w)},z=d=>{let[w,...O]=d;for(let A=0;A<O.length;A+=2){const{data:{signedUrl:se},error:te}=O[A+1];if(te)throw te;w.content=w.content.replace(O[A],se)}return w};De(async function(){_&&await r.removeChannel(_)});const W=async function(){_=await r.channel("public:guesses:game=eq."+n.code).on("postgres_changes",{event:"INSERT",schema:"public",table:"guesses",filter:"game=eq."+n.code},N).subscribe()},N=function(d){d.record.puzzle.toString()===n.puzzle&&(d.record.content==="*hint"?m.refresh():d.record.correct?t(2,s=!0):Fe.add({title:"Incorrect guess",description:d.record.content,duration:1e4,type:"error"}))},F=async function(){t(3,u=!0),await r.from("guesses").insert({content:a,puzzle:n.puzzle,game:n.code}),t(3,u=!1),t(1,a="")},R=async function(){await le("/game/"+n.code+"/"+i.next)},c=async function(){await le("/game/"+n.code)},X=["Congratulations!","Well done!","You did it!","Huzzah!","Good job!","Nice work!"];function Z(){a=this.value,t(1,a)}const P=()=>F();function J(d){l.$$.not_equal(n.code,d)&&(n.code=d,t(0,n))}function E(d){l.$$.not_equal(n.puzzle,d)&&(n.puzzle=d,t(0,n))}function ee(d){ue[d?"unshift":"push"](()=>{m=d,t(5,m)})}const x=()=>R(),Y=()=>c();return l.$$set=d=>{"params"in d&&t(0,n=d.params)},l.$$.update=()=>{l.$$.dirty&1&&n.code&&n.puzzle&&b()},[n,a,s,u,f,m,i,F,R,c,X,Z,P,J,E,ee,x,Y]}class at extends fe{constructor(e){super(),ce(this,e,nt,tt,de,{params:0})}}export{at as default};
