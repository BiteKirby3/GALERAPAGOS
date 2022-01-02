import React from "react";
import './game.css';
import { w3cwebsocket as W3CWebSocket } from "websocket";

const ROLE_FISHERMAN = "fisherman"
const ROLE_HANDYMAN = "handyman"
const ROLE_NONE = "none"

//const client = new W3CWebSocket('ws://127.0.0.1:5000');

class Game extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
        players : [],
        messages : [],
        currentPlayerID : null,
        currentMeteoImageUrl : "question_mark.gif",
        nbPlayersRest : 0,
        currentRound: 0,
        gotPlayersInfo : false,
        gameEnd : false,
    };

/*
    client.send(JSON.stringify({
        fromPage : "simulation_connexion",
      }));
      */
}

//websocket
componentDidMount() {
  /*
    client.onopen = () => {
      console.log('Client : WebSocket Client Connected');
    };
    client.onmessage = (message) => {
      var obj = JSON.parse(message.data)
      console.log(obj)
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
    };*/
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


  render() {
    return (
      <div className="home">
        <div className="scrollmenu">
        <div className="players">
          <label>Joueurs restants :</label>
          <ul className="no-bullets">
            <div className="flex-container">
                <li>
                  <p>Joueur 1</p>
                  <img src={process.env.PUBLIC_URL + "fishman.png"} width={"120px"} height={"200px"} alt="pêcheur"></img>
                  <p>Egoïsme :</p>
                  <p>Intelligence :</p>
                </li> 
            </div>
                <div className="flex-container">
                  <li>
                    <p>Joueur 2</p>
                    <img src={process.env.PUBLIC_URL + "normal_person.png"} width={"130px"} height={"200px"} alt="Météo"></img>
                    <p>Egoïsme :</p>
                    <p>Intelligence :</p>
                  </li>
                </div> 
                <div className="flex-container">
                  <li>
                    <p>Joueur 3</p>
                    <img src={process.env.PUBLIC_URL + "woodmaker.png"} width={"130px"} height={"200px"} alt="Météo"></img>
                    <p>Egoïsme :</p>
                    <p>Intelligence :</p>
                  </li>
                </div> 
                <div className="flex-container">
                <li><p>Joueur 1</p><img src={process.env.PUBLIC_URL + "fishman.png"} width={"120px"} height={"200px"} alt="pêcheur"></img></li> 
                <p>Egoïsme :</p>
                    <p>Intelligence :</p>
                </div>
                <div className="flex-container"><li><p>Joueur 2</p><img src={process.env.PUBLIC_URL + "normal_person.png"} width={"130px"} height={"200px"} alt="Météo"></img><p>Egoïsme :</p>
                    <p>Intelligence :</p></li></div> 
                <div className="flex-container"><li><p>Joueur 3</p><img src={process.env.PUBLIC_URL + "woodmaker.png"} width={"130px"} height={"200px"} alt="Météo"></img><p>Egoïsme :</p>
                    <p>Intelligence :</p></li></div> 
          </ul>
        </div>
        </div>

        <div className="gameInfo">
          <div className="log">
            <label>Log : </label>
            <textarea className="gameLog">
            </textarea>
          </div>

          <div className="currentRound">
            <p>
              <label>Tour : </label>
              <label id="roundCount">{this.state.currentRound}</label>
            </p>
          </div>

          <div className="meteo">
            <label>Météo : </label>
            <img src={process.env.PUBLIC_URL + this.state.currentMeteoImageUrl} width={"100px"} height={"100px"} alt="Météo"></img>
          </div>

          <div className="counter">
            <label>Compteur : </label>
            <table>
            <tr>
              <th>Radeau</th>
              <th>Bois</th>
              <th>Eau</th>
              <th>Poisson</th>
            </tr>
            <tr>
              <td>0</td>
              <td>0</td>
              <td>0</td>
              <td>0</td>
            </tr>
            </table> 
          </div>
        </div>
    </div>
    );
  }
}

export default Game;