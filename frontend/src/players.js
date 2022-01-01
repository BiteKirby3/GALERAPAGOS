import React from "react";
import PageDescription from "./pageDescription";
import { NavLink } from "react-router-dom";
import './players.css';
import { w3cwebsocket as W3CWebSocket } from "websocket";

const ROLE_FISHERMAN = "fisherman"
const ROLE_HANDYMAN = "handyman"
const ROLE_NONE = "none"

const client = new W3CWebSocket('ws://127.0.0.1:5000');

class Players extends React.Component {
  constructor(props) {
    super(props);
    const players = []

    this.state = {
        players : players,
        nbPlayers : 0,
        gotNbPlayers : false,
    };    
    
    this.handleChangeName = this.handleChangeName.bind(this);
    this.handleChangeSelfishness = this.handleChangeSelfishness.bind(this);
    this.handleChangeIntelligence = this.handleChangeIntelligence.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);

    client.send(JSON.stringify({
        fromPage : "players_connexion",
      }));
}

//websocket
componentWillMount() {
    client.onopen = () => {
      console.log('Client : WebSocket Client Connected');
    };
    client.onmessage = (message) => {
      var obj = JSON.parse(message.data)
      this.setState({nbPlayers : obj.nbPlayers});
      for (var i = 0; i < this.state.nbPlayers; i++) {
        this.state.players.push({
            id : i,
            name : "",
            role : ROLE_NONE,
            selfishness : 0, 
            intelligence : 0,   
        });
    }
      this.setState({gotNbPlayers : true})
    };
  }

handleSubmit(event) {
    console.log("redirect to simulation")
    client.send(JSON.stringify({
      fromPage : "players",
      players : this.state.players
    }));
  }

updatePlayer(index, attributes){
    let players = [...this.state.players];
    let player = {
        ...players[index],
        ...attributes
    }
    players[index] = player;
    this.setState({players : players});
}

handleChangeName(event, index){
    this.updatePlayer(index,{name: event.target.value})
}
handleChangeRole(event, index){
    this.updatePlayer(index,{role: event.target.value})
}
handleChangeSelfishness(event, index){
    this.updatePlayer(index,{selfishness: Number(event.target.value)})
}
handleChangeIntelligence(event, index){
    this.updatePlayer(index,{intelligence:  Number(event.target.value)})
}

  render() {
    return (
      <div className="home">
        <PageDescription url_img={process.env.PUBLIC_URL + "/actions.jpg"} page_title={"Joueurs"} descr_text={"Veuillez saisir les informations concernant les joueurs :"} />
        <div className="scrollmenu">
        {this.state.gotNbPlayers ? (
                Array.from({length: this.state.nbPlayers},(_, i) => (
                    <div key={i}>
                        Joueur {i+1} :
                        <p>
                        <label>
                            Nom : <input type="text" value={this.state.players[i].name} onChange={event => this.handleChangeName(event, i)} />     
                        </label>
                        </p>
                        <p>
                        <label>
                            Rôle : <select value={this.state.players[i].role} onChange={event => this.handleChangeRole(event, i)}>
                                <option defaultValue={ROLE_NONE}>Rien</option>
                                <option value={ROLE_FISHERMAN}>Pêcheur</option>
                                <option value={ROLE_HANDYMAN}>Bucheron</option>
                            </select>    
                        </label>
                        </p>
                        <p>
                            <label>
                                Egoïsme : <input type="number" value={this.state.players[i].selfishness}  min={0} max={10} step={1} onChange={event => this.handleChangeSelfishness(event, i)}/>   
                            </label>
                        </p>
                        <p>
                            <label>
                                Intelligence : <input type="number" value={this.state.players[i].intelligence}  min={0} max={10} step={1} onChange={event => this.handleChangeIntelligence(event, i)}/>   
                            </label>
                        </p>
                </div>
                ))     ):(<div></div>)
                }
            </div>
          <NavLink className="nav-link" to="/simulation">
              <input type="submit" value="Lancer Simulation" onClick={this.handleSubmit}/>
          </NavLink>   
    </div>
    );
  }
}

export default Players;