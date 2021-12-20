<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Beego Form</title>
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.5.0/css/bootstrap.min.css"
    integrity="sha384-9aIt2nRpC12Uk9gS9baDl411NQApFmC26EwAOH8WgZl5MYYxFfc+NcPb1dKGj7Sk" crossorigin="anonymous">
    <link href="https://fonts.googleapis.com/css2?family=Open+Sans:wght@300&display=swap" rel="stylesheet">
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css">
    <link rel="stylesheet" href="//code.jquery.com/ui/1.12.1/themes/base/jquery-ui.css">
</head>
<body>
<div class="mx-4" style="margin:auto; border: 5px solid gray;">
  <p style="text-align:center; font-size:40px;">Enter your details:</p>
  <form action="/" method="post" id="user" novalidate>
  <div class="row  m-4">
    <div class="col">
      <input type="text" class="form-control" placeholder="First name" name="fname" required>
      {{ if .Fname }}
      {{.flash.error}}
      {{end}}
    </div>
    <div class="col">
      <input type="text" class="form-control" placeholder="Last name" name="lname" required>
      {{if .Lname}}
      {{.flash.error}}
      {{end}}
    </div>
  </div>
  <div class="row  m-4">
      <div class="col">
      <input type="password" class="form-control" placeholder="Password" name="password" required>
      </div>
      <div class="col">
      <input type="text" class="form-control" placeholder="8801XXXXXXX" name="phone" required>
      {{if .Phone}} {{.flash.warning}}  {{end}}
      </div>
  </div>
  <div class="row  m-4">
      <div class="col">
      <input type="text" class="form-control" placeholder="Email" name="email" required>
      {{if .Email}}
      {{.flash.notice}}
      {{end}}
      </div>
      <div class="col">
      <input type="text" class="form-control" placeholder="Date of Birth(mm/dd/yyyy)" name="dob" required>
      </div>
  </div>
    <div class="row  m-4" style="text-align:center;">
      <div class="col">
      <input type="submit" class="form-control bg-info" style="font-size:20px; color:white;" value="Submit">
      {{if .Success}}
      {{.flash.success}}
      {{end}}
      </div>

  </div>
  </div>
  </form>
</div>
</body>
</html>