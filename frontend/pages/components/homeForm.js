import React from "react";

class HomeForm extends React.Component {
    constructor(props) {
      super(props);
      this.state = {      
        nbPlayers: 2,
        nbTurns : 2,   
      };
      this.handleChangeNbPlayers = this.handleChangeNbPlayers.bind(this);
      this.handleChangeNbTurns = this.handleChangeNbTurns.bind(this);
      this.handleSubmit = this.handleSubmit.bind(this);
    }
  
    handleChangeNbPlayers(event) {    this.setState({nbPlayers: event.target.value});  }
    handleChangeNbTurns(event) {    this.setState({nbTurns: event.target.value});  }
    handleSubmit(event) {
      alert('Nombre de joueurs : ' + this.state.nbPlayers+"\n"+'Nombre de tours : ' + this.state.nbTurns);
      event.preventDefault();
    }
  
    render() {
      return (
        <form onSubmit={this.handleSubmit}>
          <p>
          <label>
            Nombre de joueurs : <input type="number" value={this.state.nbPlayers} onChange={this.handleChangeNbPlayers} />     
          </label>
          </p>
          <p>
          <label>
            Nombre de tours : <input type="number" value={this.state.nbTurns} onChange={this.handleChangeNbTurns} />    
          </label>
          </p>
          <input type="submit" value="Suivant" />
        </form>
      );
    }
  }

  export default HomeForm;