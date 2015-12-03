var Root = React.createClass({
  mixins: [React.addons.LinkedStateMixin],
  getInitialState: function() {
    return {
      users: [],
	  firstName: "",
	  lastName: ""
    };
  },
  componentDidMount: function() {
      var path = this.getPath();
      var account = this.getAccount(path);
      console.log(account);
    var u = "/api/accounts/" + account + "/users"
	  $.getJSON(u, function(data) {
      this.setState({users: data});
    }.bind(this));
  },
 getUrl: function() {
    return '' + window.location;
  },
  getPath: function() {
    return window.location.pathname + window.location.search;
  },
  getAccount: function(path) {
    var urlPieces = path.split("/");
    var account = (urlPieces[2]) ? urlPieces[2] : "";
    return account;
  },
  handleSubmit: function(e) {
	e.preventDefault();
      var path = this.getPath();
      var account = this.getAccount(path);
    var u = "/api/accounts/" + account + "/users"
    $.ajax({
	url: u,
      data: JSON.stringify({"first_name": this.state.firstName, "last_name": this.state.lastName}),
      method: "POST",
      success: function(data, textStatus, request) {
		  var newNames = this.state.users;
		  var href = request.getResponseHeader('Location');
		   newNames.push({"href": href, "id": href.substr(href.length -1), "first_name": this.state.firstName, "last_name": this.state.lastName});
        this.setState({
			firstName: "",
			lastName: "",
          users: newNames
		});
	  }.bind(this),
	  error: function(request, textStatus, response) {
		  var resp = JSON.parse(request.responseText);
         alert(resp[0].title +": " +  resp[0].msg);
        },
    });
  },
  render: function() {

    var listElts = [];
    for(var i = 0; i < this.state.users.length; i++) {
      var user = this.state.users[i];
      listElts.push(<li key={user.id} className="list-group-item">{user.href} - {user.first_name} {user.last_name}</li>);
    }
    return (
      <div className="container users">
        <div className="row">
          <ul className="list-group">
            {listElts}
          </ul>
        </div>
        <div className="row">
          <h3>Create New User</h3>
          <form onSubmit={this.handleSubmit}>
            <div className="form-group">
              <label htmlFor="accountName">User Name</label>
			  <input type="text" className="form-control" id="firstName" placeholder="First Name" valueLink={this.linkState('firstName')}/> 
			  <input type="text" className="form-control" id="lastName" placeholder="Last Name"  valueLink={this.linkState('lastName')}/>

            </div>
            <button className="btn btn-default">Create</button>
          </form>
        </div>
      </div>
    );
  },
});

ReactDOM.render(<Root/>, document.getElementById("page-content"));
