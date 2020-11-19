<!DOCTYPE html>

<html>
<head>
    <title>用户登录</title>
    <link href="http://cdn.static.runoob.com/libs/bootstrap/3.3.7/css/bootstrap.min.css" rel="stylesheet">
    <script type="text/javascript" src="http://cdn.static.runoob.com/libs/jquery/2.1.1/jquery.min.js"></script>
    <script src="http://cdn.static.runoob.com/libs/bootstrap/3.3.7/js/bootstrap.min.js"></script>
  </head>
  <body>
 <form id="user" action="http://localhost:8082" method="post">
用户名：<input name="Name" type="text"/>
密码：<input name="Password1" type="text"/>
   <input id="Submit1" type="submit" value="登录"onclick="login()"  />   
   <span class="error"></span>
  </form>

    <script type="text/javascript">
        function login(){
            $.ajax({
                type:'post',
                url:'/test/log',
                data:{
                    "name":$("#name").val(),
                    "password":$("#password").val()
                },
                success:function(result){
                    if(result.State > 0){
                        alert(result.Message)
                    }else {
                        alert("登录成功");
                      location.href = "/"
                    }
                }
            })
        }
            function sign() {
                      location.href = "/test/sign"
            }
        </body>
    </script>
</html>