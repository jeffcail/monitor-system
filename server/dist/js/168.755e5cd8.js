"use strict";(self["webpackChunkweb"]=self["webpackChunkweb"]||[]).push([[168],{6168:function(e,a,l){l.r(a),l.d(a,{default:function(){return A}});l(9737);var t=l(1870),i=l(6810),u=l(4195),n=l(971),d=l(366),p=l(8276),o=l(4181),s=l(3378),m=l(6252),r=l(3577),c=l(2262),_=l(2201),g=l(842),v=l(1348);const w=e=>((0,m.dD)("data-v-1d329180"),e=e(),(0,m.Cn)(),e),h={class:"machine-list"},f={style:{display:"flex","align-items":"center"}},k={style:{"margin-left":"10px"}},z={style:{display:"flex","align-items":"center"}},W={style:{"margin-left":"10px"}},y={style:{display:"flex","align-items":"center"}},b={style:{"margin-left":"10px"}},U={style:{display:"flex","align-items":"center"}},x={style:{"margin-left":"10px"}},C=(0,m.Uk)("发送指令"),H=(0,m.Uk)("升级客户端"),V={class:"demo-pagination-block"},P=(0,m.Uk)(" 指令    "),B={class:"dialog-footer"},D=(0,m.Uk)("取消"),I=(0,m.Uk)("发送"),S=w((()=>(0,m._)("div",{class:"el-upload__text"},[(0,m.Uk)(" 拖到此区域上传 或者 "),(0,m._)("em",null,"点击上传")],-1))),j=w((()=>(0,m._)("div",{class:"el-upload__tip",style:{color:"red"},align:"left"},[(0,m._)("span",null,"友情提示:"),(0,m.Uk)(),(0,m._)("br"),(0,m._)("span",null,"1. 包名不能带后缀"),(0,m.Uk)(),(0,m._)("br"),(0,m._)("span",null,"2. 包文件大小不能超过100M"),(0,m.Uk)(),(0,m._)("br"),(0,m._)("span",null,"3. 版本号不得重复"),(0,m.Uk)(),(0,m._)("br"),(0,m._)("span",null,"4. 当前版本号不得小于上次升级版本号")],-1))),Y=(0,m.Uk)(" 输入版本号:     "),Z={style:{"margin-left":"-400px"}},$=(0,m.Uk)("取消"),q=(0,m.Uk)("上传"),E=w((()=>(0,m._)("div",{style:{"marigin-top":"10px","margin-bottom":"10px"},align:"left"}," 记录 ",-1))),F={class:"demo-pagination-block"};var L={__name:"machine",setup(e){(0,c.iH)(null);const a=(0,c.iH)(!1),l=e=>{a.value=!0,w.value.machine_code=e.machine_code,w.value.machine_ip=e.ip,w.value.machine_hostname=e.host_name,w.value.machine_remark=e.remark,G.value.machine_ip=e.ip,J()},w=(0,c.iH)({go_file:"",machine_code:"",machine_ip:"",machine_hostname:"",machine_remark:"",upgrade_version:""}),L=e=>{let a=e.file;w.value.go_file=a,M(1)},M=async e=>{if(1===e);else{let e=new FormData;e.append("go_file",w.value.go_file),e.append("machine_code",w.value.machine_code),e.append("machine_ip",w.value.machine_ip),e.append("machine_hostname",w.value.machine_hostname),e.append("machine_remark",w.value.machine_remark),e.append("upgrade_version",w.value.upgrade_version);let l=await(0,g.xk)(e);l&&2e3===l.code&&(v.z8.success(l.msg),a.value=!1)}},R=(0,c.iH)([]),A=(0,c.iH)(0),G=(0,c.iH)({page:1,page_size:10,machine_ip:""}),J=async()=>{let e={page:G.value.page,page_size:G.value.page_size,machine_ip:G.value.machine_ip},a=await(0,g._3)(e);2e3!==a.code&&(a.data=[]),R.value=a.data.list,A.value=a.data.total},K=()=>{G.value.page_size=row,G.value.page=1,J()},N=()=>{G.value.page=row,J()},O=async e=>{let a={machine_code:e.machine_code,ip:e.ip,remark:e.remark},l=await(0,g.gS)(a);console.log(l),2e3===l.code&&(v.z8.success(l.msg),de())},Q=(0,c.iH)(!1),T=(0,c.iH)(!1),X=(0,c.iH)(!1),ee=((0,_.tv)(),(0,c.iH)(!1)),ae=(0,c.iH)({content:"",ip:""}),le=e=>{ee.value=!0,ae.value.ip=e.ip},te=async()=>{let e={ip:ae.value.ip,content:ae.value.content},a=await(0,g.$p)(e);2e3===a.code&&(v.z8.success("指令发送成功"),ee.value=!1,de())},ie=(0,c.iH)(0),ue=(0,c.iH)([]),ne=(0,c.iH)({page:1,page_size:10}),de=((0,c.iH)(),async()=>{let e={page:ne.value.page,page_size:ne.value.page_size},a=await(0,g.cB)(e);2e3!==a.code&&(a.data=[]),ue.value=a.data.list,ie.value=a.data.total});de();const pe=e=>{ne.value.page_size=e,ne.value.page=1,de()},oe=e=>{ne.value.page=e,de()};return(e,c)=>{const _=d.$Y,g=s.EZ,v=(0,m.up)("timer"),J=o.gn,de=p.mi,se=d.eI,me=n.R,re=u.nH,ce=u.ly,_e=i.d0,ge=(0,m.up)("upload-filled"),ve=t.LW;return(0,m.wg)(),(0,m.iD)(m.HY,null,[(0,m._)("div",h,[(0,m.Wm)(se,{data:ue.value,style:{width:"100%"}},{default:(0,m.w5)((()=>[(0,m.Wm)(_,{label:"机器码",width:"290"},{default:(0,m.w5)((e=>[(0,m._)("div",f,[(0,m._)("span",k,(0,r.zw)(e.row.machine_code),1)])])),_:1}),(0,m.Wm)(_,{label:"IP",width:"180"},{default:(0,m.w5)((e=>[(0,m._)("div",z,[(0,m._)("span",W,(0,r.zw)(e.row.ip),1)])])),_:1}),(0,m.Wm)(_,{label:"主机名字",width:"280"},{default:(0,m.w5)((e=>[(0,m._)("div",y,[(0,m._)("span",b,(0,r.zw)(e.row.host_name),1)])])),_:1}),(0,m.Wm)(_,{label:"备注",width:"280"},{default:(0,m.w5)((e=>[(0,m.Wm)(g,{modelValue:e.row.remark,"onUpdate:modelValue":a=>e.row.remark=a,class:"w-50 m-2",size:"large",placeholder:"请输入备注",onBlur:a=>O(e.row)},null,8,["modelValue","onUpdate:modelValue","onBlur"])])),_:1}),(0,m.Wm)(_,{label:"创建时间",width:"280"},{default:(0,m.w5)((e=>[(0,m._)("div",U,[(0,m.Wm)(J,null,{default:(0,m.w5)((()=>[(0,m.Wm)(v)])),_:1}),(0,m._)("span",x,(0,r.zw)(e.row.created_at),1)])])),_:1}),(0,m.Wm)(_,{label:"操作"},{default:(0,m.w5)((({row:e})=>[(0,m.Wm)(de,{size:"small",type:"success",onClick:a=>le(e)},{default:(0,m.w5)((()=>[C])),_:2},1032,["onClick"]),(0,m.Wm)(de,{size:"small",type:"success",onClick:a=>l(e)},{default:(0,m.w5)((()=>[H])),_:2},1032,["onClick"])])),_:1})])),_:1},8,["data"]),(0,m._)("div",V,[(0,m.Wm)(me,{currentPage:ne.value.page,"onUpdate:currentPage":c[0]||(c[0]=e=>ne.value.page=e),"page-size":ne.value.page_size,"onUpdate:page-size":c[1]||(c[1]=e=>ne.value.page_size=e),"page-sizes":[10,20,30,40,50],small:Q.value,disabled:T.value,background:X.value,layout:"total, sizes, prev, pager, next, jumper",total:ie.value,"current-page":ne.value.page,onSizeChange:pe,onCurrentChange:oe},null,8,["currentPage","page-size","small","disabled","background","total","current-page"])])]),(0,m.Wm)(_e,{modelValue:ee.value,"onUpdate:modelValue":c[4]||(c[4]=e=>ee.value=e),title:"发送指令",width:"30%",draggable:""},{footer:(0,m.w5)((()=>[(0,m._)("span",B,[(0,m.Wm)(de,{onClick:c[3]||(c[3]=e=>ee.value=!1)},{default:(0,m.w5)((()=>[D])),_:1}),(0,m.Wm)(de,{type:"primary",onClick:te},{default:(0,m.w5)((()=>[I])),_:1})])])),default:(0,m.w5)((()=>[(0,m.Wm)(ce,{model:ae.value},{default:(0,m.w5)((()=>[(0,m.Wm)(re,null,{default:(0,m.w5)((()=>[P,(0,m.Wm)(g,{modelValue:ae.value.content,"onUpdate:modelValue":c[2]||(c[2]=e=>ae.value.content=e),autocomplete:"off",placeholder:"指令为linux命令 如: ls && mkdir && cd /root/xxx/stat.sh",style:{width:"400px"}},null,8,["modelValue"])])),_:1})])),_:1},8,["model"])])),_:1},8,["modelValue"]),(0,m.Wm)(_e,{modelValue:a.value,"onUpdate:modelValue":c[10]||(c[10]=e=>a.value=e),title:"升级客户端",width:"40%",draggable:""},{default:(0,m.w5)((()=>[(0,m.Wm)(ce,{model:w.value},{default:(0,m.w5)((()=>[(0,m.Wm)(ve,{class:"upload-demo",drag:"",multiple:"","http-request":L,limit:1},{tip:(0,m.w5)((()=>[j])),default:(0,m.w5)((()=>[(0,m.Wm)(J,{class:"el-icon--upload"},{default:(0,m.w5)((()=>[(0,m.Wm)(ge)])),_:1}),S])),_:1}),(0,m.Wm)(re,null,{default:(0,m.w5)((()=>[Y,(0,m.Wm)(g,{modelValue:w.value.upgrade_version,"onUpdate:modelValue":c[5]||(c[5]=e=>w.value.upgrade_version=e),autocomplete:"off",style:{width:"200px"}},null,8,["modelValue"])])),_:1})])),_:1},8,["model"]),(0,m._)("div",Z,[(0,m.Wm)(de,{onClick:c[6]||(c[6]=e=>a.value=!1)},{default:(0,m.w5)((()=>[$])),_:1}),(0,m.Wm)(de,{type:"primary",onClick:c[7]||(c[7]=e=>M(2))},{default:(0,m.w5)((()=>[q])),_:1})]),E,(0,m.Wm)(se,{data:R.value,style:{width:"100%"}},{default:(0,m.w5)((()=>[(0,m.Wm)(_,{prop:"created_at",label:"升级时间",width:"200"}),(0,m.Wm)(_,{prop:"package_name",label:"包名",width:"200"}),(0,m.Wm)(_,{prop:"upgrade_version",label:"版本号"})])),_:1},8,["data"]),(0,m._)("div",F,[(0,m.Wm)(me,{currentPage:G.value.page,"onUpdate:currentPage":c[8]||(c[8]=e=>G.value.page=e),"page-size":G.value.page_size,"onUpdate:page-size":c[9]||(c[9]=e=>G.value.page_size=e),"page-sizes":[10,20,30,40,50],small:Q.value,disabled:T.value,background:X.value,layout:"total, sizes, prev, pager, next, jumper",total:A.value,"current-page":G.value.page,onSizeChange:K,onCurrentChange:N},null,8,["currentPage","page-size","small","disabled","background","total","current-page"])])])),_:1},8,["modelValue"])],64)}}},M=l(3744);const R=(0,M.Z)(L,[["__scopeId","data-v-1d329180"]]);var A=R}}]);
//# sourceMappingURL=168.755e5cd8.js.map