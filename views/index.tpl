<!DOCTYPE html>

<html>
<head>
  <title>用户论坛</title>
</head>

<body>
Welcome to the forum.
<form id="user" action="http://localhost:8082" >

<!--
用户名：<input name="Name" type="text"/>
密码：<input name="Password" type="text"/>
-->

    <div class="form-group">
        <input type="button" value="注册" onclick=javascrtpt:sign()>
        <input type="button" value="登录" onclick=javascrtpt:log()>
    </div>
<script>
    function sign() {
    window.location.href="http://localhost:8082/test/sign";
    }
 </script>
 <script>
    function log() {
    window.location.href="http://localhost:8082/test/log";
    }
</script>
<script src="/static/js/reload.min.js"></script>
<script type="text/javascript" src="http://cdn.static.runoob.com/libs/jquery/2.1.1/jquery.min.js"></script>

</body>
</html>
