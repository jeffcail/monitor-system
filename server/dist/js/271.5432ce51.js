"use strict";(self["webpackChunkweb"]=self["webpackChunkweb"]||[]).push([[271],{1998:function(e,a,l){l.r(a),l.d(a,{default:function(){return k}});var t=l(971),s=l(366),i=l(4181),n=l(6252),d=l(3577),u=l(2262),r=l(842);const p=(0,n._)("div",{class:"add-admin"},null,-1),g={class:"admin-list"},o={style:{display:"flex","align-items":"center"}},c={style:{"margin-left":"10px"}},m={style:{display:"flex","align-items":"center"}},_={style:{"margin-left":"10px"}},w={style:{display:"flex","align-items":"center"}},v={style:{"margin-left":"10px"}},f={style:{display:"flex","align-items":"center"}},y={style:{"margin-left":"10px"}},z={style:{display:"flex","align-items":"center"}},b={style:{"margin-left":"10px"}},h={class:"demo-pagination-block"};var x={__name:"log",setup(e){const a=(0,u.iH)(0),l=(0,u.iH)([]),x=(0,u.iH)({page:1,page_size:10}),W=async()=>{let e={page:x.value.page,page_size:x.value.page_size},t=await(0,r.N1)(e);console.log(t),2e3!==t.code&&(t.data=[]),l.value=t.data.list,a.value=t.data.total};W();const k=e=>{x.value.page_size=e,x.value.page=1,W()},C=e=>{x.value.page=e,W()};return(e,u)=>{const r=s.$Y,W=(0,n.up)("timer"),H=i.gn,P=s.eI,D=t.R;return(0,n.wg)(),(0,n.iD)(n.HY,null,[p,(0,n._)("div",g,[(0,n.Wm)(P,{data:l.value,style:{width:"100%"}},{default:(0,n.w5)((()=>[(0,n.Wm)(r,{label:"ID",width:"60"},{default:(0,n.w5)((e=>[(0,n._)("div",o,[(0,n._)("span",c,(0,d.zw)(e.row.id),1)])])),_:1}),(0,n.Wm)(r,{label:"用户名",width:"80"},{default:(0,n.w5)((e=>[(0,n._)("div",m,[(0,n._)("span",_,(0,d.zw)(e.row.admin_username),1)])])),_:1}),(0,n.Wm)(r,{label:"路径",width:"180"},{default:(0,n.w5)((e=>[(0,n._)("div",w,[(0,n._)("span",v,(0,d.zw)(e.row.url),1)])])),_:1}),(0,n.Wm)(r,{label:"日志",width:"680"},{default:(0,n.w5)((e=>[(0,n._)("div",f,[(0,n._)("span",y,(0,d.zw)(e.row.content),1)])])),_:1}),(0,n.Wm)(r,{label:"操作时间",width:"180"},{default:(0,n.w5)((e=>[(0,n._)("div",z,[(0,n.Wm)(H,null,{default:(0,n.w5)((()=>[(0,n.Wm)(W)])),_:1}),(0,n._)("span",b,(0,d.zw)(e.row.created_at),1)])])),_:1}),(0,n.Wm)(r)])),_:1},8,["data"]),(0,n._)("div",h,[(0,n.Wm)(D,{currentPage:x.value.page,"onUpdate:currentPage":u[0]||(u[0]=e=>x.value.page=e),"page-size":x.value.page_size,"onUpdate:page-size":u[1]||(u[1]=e=>x.value.page_size=e),"page-sizes":[10,20,30,40,50],small:e.small,disabled:e.disabled,background:e.background,layout:"total, sizes, prev, pager, next, jumper",total:a.value,"current-page":x.value.page,onSizeChange:k,onCurrentChange:C},null,8,["currentPage","page-size","small","disabled","background","total","current-page"])])])],64)}}};const W=x;var k=W}}]);
//# sourceMappingURL=271.5432ce51.js.map