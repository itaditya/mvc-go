<!DOCTYPE html>
<html>

<head>
  <title>Beego</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <style>
    .users-table {
      text-align: left;
    }
  </style>
</head>
<body>
  <h1>User Dashboard</h1>
  <hr />
  <h4>Add User</h4>
  <form class="js-user-add-form" method="POST" action="/users">
    <label>
      name：<input name="user_name" type="text" required />
    </label>
    <label>
      email：<input name="email" type="text" required />
    </label>
    <label>
      password : <input name="password" type="password" required />
    </label>
    <input type="submit" value="submit" />
  </form>
  <br />
  <h4>All Users</h4>
  <table class="users-table js-users-table">
    <thead>
      <tr>
        <th>Id</th>
        <th>Name</th>
        <th>Email</th>
        <th>Action</th>
      </tr>
    </thead>
    <tbody class="js-table-body">
      {{range .Users}}
        <tr>
          <td>{{.UserId}}</td>
          <td>{{.UserName}}</td>
          <td>{{.Email}}</td>
          <td><button class="js-delete-user-button" data-id="{{.UserId}}">Delete</button></td>
        </tr>
      {{end}}
    </tbody>
  </table>
  <script src="/static/js/reload.min.js"></script>
  <script src="/static/js/main.js"></script>
</body>

</html>
