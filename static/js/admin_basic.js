$(".J_menuItem").on('click', function(e){
    e.preventDefault();
    var url = $(this).attr('href');
    if (url.indexOf("?") > 0){
        var proccess_url = url.substring(0,url.indexOf("?"));
    }else{
        proccess_url = url
    }
// 如果点击首页跳转整个页面
    if (url.search('/admin/index') >= 0){
        if (top.location.href.search("/admin/index")< 1){
            top.location = "/admin/index"
        }
        return
    }
    var link_content = $(this).find("span").html();
    if(link_content == undefined || link_content == ""){
        //无图标
        var link_content = $(this).html();
    }
    var url_array = proccess_url.split("/")
    url_array.shift();
    var tab_id = url_array.join("_");
    var tab_li_template = $(".tab-template .tab-li").html();
    var tab_li = tab_li_template.replace(/\{\id}/g,tab_id).replace("{link_content}",link_content);
    var tab_content_template = $(".tab-template .tab-content-container").html();
    var tab_content = tab_content_template.replace("#{url}",url).replace("{id}",tab_id).replace("{name}",url.replace(/\//g,"/"));
    if($(".content-tab a[href='#"+ tab_id +"']").length < 1){
        //插入内容
        $(".content-tab ul.nav-tabs").append(tab_li);
        $(".content-tab div.tab-content").append(tab_content);
    }
    //激活新tab
    $(".content-tab a[href='#"+ tab_id +"']").tab("show");
})

$(".content-tab").on("click",".close-tab",function(){
    //关闭tab
    //找到tab_id
      var tab_id = $(this).parent().attr("href");
    //找到对应的tab-pane
      var tab_pane = $(".tab-content " + tab_id);
     //当前tab
      var tab_li = $(this).parent().parent();
    //需要激活的上一个tab
     var pre_tab = $(tab_li).prev();
    //移除tab-pane tab_li
     tab_pane.detach();
     tab_li.detach();
    //激活上一个tab
     pre_tab.find("a").tab("show");
  })