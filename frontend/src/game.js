import React from "react";
import './game.css';
import { w3cwebsocket as W3CWebSocket } from "websocket";

const ROLE_FISHERMAN = "fisherman"
const ROLE_HANDYMAN = "handyman"
const ROLE_NONE = "none"

const client = new W3CWebSocket('ws://127.0.0.1:5000');

class Game extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
        players : [],
        messages : [],
        currentPlayerID : null,
        nbPlayersRest : 0,
        gotPlayersInfo : false,
        gameEnd : false,
    };

    this.handleChangeName = this.handleChangeName.bind(this);
    this.handleChangeSelfishness = this.handleChangeSelfishness.bind(this);
    this.handleChangeIntelligence = this.handleChangeIntelligence.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);

    client.send(JSON.stringify({
        fromPage : "simulation_connexion",
      }));
}

//websocket
componentWillMount() {
    client.onopen = () => {
      console.log('Client : WebSocket Client Connected');
    };
    client.onmessage = (message) => {
      var obj = JSON.parse(message.data)
      if (obj.messageType == "gameStart"){
        this.setState({gotPlayersInfo : true})
      } else if (obj.messageType == "roundStart"){

      } else if (obj.messageType == "info"){

      } else if (obj.messageType == "meteo"){

      } else if (obj.messageType == "action"){

      } else if (obj.messageType == "death"){

      } else if (obj.messageType == "roundEnd"){

      } else if (obj.messageType == "constructRaft"){

      } else if (obj.messageType == "gameEnd"){
        this.setState({ gameEnd : true})
      }
      /*
      this.setState({nbPlayers : obj.nbPlayers});
      for (var i = 0; i < this.state.nbPlayers; i++) {
        this.state.players.push({
            id : i,
            name : "",
            role : ROLE_NONE,
            selfishness : 0,
            intelligence : 0,
        });
    }*/
    };
  }
  componentDidMount() {
    client.onopen = () => {
      console.log('Client : WebSocket Client Connected');
    };
    client.onmessage = (message) => {
      console.log(message);
      const data = JSON.parse(message.data)
      console.log(data);
    };
    client.send(JSON.stringify({
      fromPage : "simulation_connexion",
    }));
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
      <div className="game">
        {this.state.gotPlayersInfo ?
        (
        <div className="gameLog">
        <textarea>
        </textarea>
        </div>
        ):
        (
        <div className="unload">
          <h1>Veuillez attendre ... </h1>
          <img
                    className="page_image"
                    src="../public/loading.gif"
                    alt=""
                />
        </div>)
        }
        <div className="playerCard">
        {this.state.gotPlayersInfo ? (
                Array.from({length: this.state.nbPlayersRest},(_, i) => (
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

    </div>
    );
  }
}

export default Game;