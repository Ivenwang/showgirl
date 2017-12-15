<!DOCTYPE html>
<html>
  <head>
    <title>ELKWatchDog</title>

    <link rel='stylesheet' href='/static/css/bootstrap.css' />
    <link href="/static/css/bootstrap-responsive.css" rel="stylesheet">
<!--

    <link href="http://libs.baidu.com/bootstrap/3.0.3/css/bootstrap.min.css" rel="stylesheet">
    <script src="http://libs.baidu.com/jquery/2.0.0/jquery.min.js"></script>
    <script src="http://libs.baidu.com/bootstrap/3.0.3/js/bootstrap.min.js"></script>
-->    

    <style type="text/css">
      body {
        padding-top: 60px;
        padding-bottom: 40px;
      }
    </style>

  </head>
  <body>
  
    <div class="navbar navbar-fixed-top">
      <div class="navbar-inner">
        <div class="container">
          <a class="btn btn-navbar" data-toggle="collapse" data-target=".nav-collapse">
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
          </a>
          <a class="brand" href="/">Dashboard For WatchDog</a>
          <div class="nav-collapse">
            <ul class="nav">
              <li class="active"><a href="/admin" target="_blank">首页</a></li>
              <li>
                <form role="form" class="form-inline">
                  <input type="text" class="form-control" id="keyword" placeholder="请输入模块名">
                  <button type="submit" class="btn btn-default" id="search-button">搜索</button>
                </form>
              </li>
              <li>
                <button type="submit" class="btn btn-default" id="add-button">添加监控项</button>
              </li>
            </ul>
          </div>
        </div>
      </div>
    </div>

    <div id="container" class="container">
      <div id="watchregister"></div>
      {{.LayoutContent}}
      <hr />
      <p/>
      <footer>
        <p>金山云VCLOUD</p>
      </footer>
    </div>
  </body>

  <script src="/static/js/jquery.js"></script>
  <script src="/static/js/bootstrap.js"></script>
  <script type="text/javascript">
    $(document).ready(function(){
      $("#search-button").click(function(){
        var url = "/admin?k="+$("#keyword").val()
        window.open(url)
      });
      $("#add-button").click(function(){
        var url = "/admin_reg"
        window.open(url)
      });
    });
  </script>
</html>
