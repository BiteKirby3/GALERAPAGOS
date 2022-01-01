import React from "react";
import PageDescription from "./pageDescription";
import { NavLink } from "react-router-dom";
import { w3cwebsocket as W3CWebSocket } from "websocket";

const client = new W3CWebSocket('ws://127.0.0.1:5000');

class Home extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      nbPlayers: 3,
      nbTurns : 2,         
    };
    this.handleChangeNbPlayers = this.handleChangeNbPlayers.bind(this);
    this.handleChangeNbTurns = this.handleChangeNbTurns.bind(this);
    this.handleClick = this.handleClick.bind(this);
  }

  //websocket
  componentWillMount() {
    client.onopen = () => {
      console.log('Client : WebSocket Client Connected');
    };
    client.onmessage = (message) => {
      //console.log(message);
    };
  }
  

  handleChangeNbPlayers(event) { this.setState({nbPlayers: event.target.value});  }
  handleChangeNbTurns(event) { this.setState({nbTurns: event.target.value});  }
  handleClick(event) {
    console.log("redirect to players")
    client.send(JSON.stringify({
      fromPage:"home",
      nbPlayers: ""+this.state.nbPlayers,
      nbTurns: ""+this.state.nbTurns
    }));
  }



  render() {
    return (
      <div className="home">
        <PageDescription url_img={process.env.PUBLIC_URL + "/player.jpg"} page_title={"Après le naufrage de votre bateau..."} descr_text={"    votre groupe de survivants se retrouve sur une île déserte où l’eau \net la nourriture se font rares. Seule solution pour échapper à ce cauchemar : \nconstruire ensemble un grand radeau pour embarquer les survivants, \nmais le temps presse car un ouragan pointe à l’horizon..."} />
          <p>
          <label>
            Nombre de joueurs : <input type="number" step={1} min={3} max={12} value={this.state.nbPlayers} onChange={this.handleChangeNbPlayers} />&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;   
          </label>
          <label>
            Nombre de tours : <input type="number" step={1} min={2} max={20} value={this.state.nbTurns} onChange={this.handleChangeNbTurns} />    
          </label>
          <NavLink className="nav-link" to={{ pathname:'/players', state: {nbPlayers: this.state.nbPlayers, nbTurns:this.state.nbTurns} }}>
              <input type="submit" value="Suivant" onClick={this.handleClick}/>
          </NavLink> 
          </p>
    </div>
    );
  }
}

export default Home;