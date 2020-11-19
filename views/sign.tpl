<!DOCTYPE html>

<html>
<head>
  <title>用户注册</title>
      <link href="http://cdn.static.runoob.com/libs/bootstrap/3.3.7/css/bootstrap.min.css" rel="stylesheet">
      <script type="text/javascript" src="http://cdn.static.runoob.com/libs/jquery/2.1.1/jquery.min.js"></script>
      <script src="http://cdn.static.runoob.com/libs/bootstrap/3.3.7/js/bootstrap.min.js"></script>

</head>
<body>
 <form id="test" action="http://localhost:8082" method="post">
用户名：<input name="Name" id="name" type="text"/>
密码：<input name="Password1" id="password" type="text"/>
   <input id="Submit1" type="submit" value="提交"onclick="sign()"  />   
   <input id="Reset1" type="reset" value="重置" />
   <span class="error"></span>
  </form>



</div>
<!--JS部分-->
<script type="text/javascript">
    function sign() {
        $.ajax({
            type: 'post',
            url: '/test/sign',
            data: {
                "name": $("#name").val(),
                "password": $("#password").val()
            },
            success: function (result) {
                if(result.State > 0){
                    alert(result.Message);
                }else{
                    alert("注册成功");
                location.href = "/test/sign"
                }
            }
        })
    }
</script>
</body>
</html>