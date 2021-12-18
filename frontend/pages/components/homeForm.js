import React from "react";

class HomeForm extends React.Component {
    constructor(props) {
      super(props);
      this.state = {
        nbPlayers: 3,
        nbTurns : 2,         
      };
      this.handleChangeNbPlayers = this.handleChangeNbPlayers.bind(this);
      this.handleChangeNbTurns = this.handleChangeNbTurns.bind(this);
      this.handleSubmit = this.handleSubmit.bind(this);
    }
  
    handleChangeNbPlayers(event) {    this.setState({nbPlayers: event.target.value});  }
    handleChangeNbTurns(event) {    this.setState({nbTurns: event.target.value});  }
    handleSubmit(event) {
      event.preventDefault();
      this.props.onSubmit(this.state.nbPlayers,this.state.nbTurns)
    }
  
    render() {
      return (
        <form onSubmit={this.handleSubmit}>
          <p>
          <label>
            Nombre de joueurs : <input type="number" step={1} min={3} max={12} value={this.state.nbPlayers} onChange={this.handleChangeNbPlayers} />     
          </label>
          </p>
          <p>
          <label>
            Nombre de tours : <input type="number" step={1} min={2} max={20} value={this.state.nbTurns} onChange={this.handleChangeNbTurns} />    
          </label>
          </p>
          <input type="submit" value="Suivant" />
        </form>
      );
    }
  }

  export default HomeForm;