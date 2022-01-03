import React from "react";
import './game.css';
import PageDescription from "./pageDescription";
import PlayerImage from "./playerImage";
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
        currentMeteoImageUrl : "question_mark.gif",
        nbPlayersRest : 0,
        currentRound: 0,
        gameStart : false,
        gameEnd : false,
        stockBois : 0,
        stockEau : 0,
        StockNourriture : 0,
        PlaceRadeau :0 
    };


    client.send(JSON.stringify({
        fromPage : "simulation_connexion",
    }));
    
}

//websocket
componentDidMount() {
    client.onopen = () => {
      console.log('Client : WebSocket Client Connected');
    };
    client.onmessage = (message) => {
      var obj = JSON.parse(message.data)
      if(obj.TypeEvent !== "empty"){
          console.log(obj)
          this.setState({ messages : this.state.messages.concat((obj.Description)),
            stockBois : obj.Plateau.StockBois,
            stockEau : obj.Plateau.StockEau,
            StockNourriture : obj.Plateau.StockNourriture,
            PlaceRadeau :obj.Plateau.PlaceRadeau,
            currentRound : obj.Plateau.NbTour
        })
      }

      
      if (obj.TypeEvent === "gameStart"){
        for (var i = 0; i <obj.ListJoueurs.length; i++) {
          var role_joueur
          if(obj.ListJoueurs[i].Bucheron){
            role_joueur=ROLE_HANDYMAN
          } else if(obj.ListJoueurs[i].Pecheur){
            role_joueur=ROLE_FISHERMAN
          } else {
            role_joueur=ROLE_NONE
          }
          this.state.players.push({
              id : obj.ListJoueurs[i].ID,
              name : obj.ListJoueurs[i].Nom,
              role : role_joueur,
              estDead : false,
              selfishness : obj.ListJoueurs[i].Egoisme, 
              intelligence : obj.ListJoueurs[i].Intelligence,   
          });
      }
        this.setState({gameStart : true})
      } else if (obj.TypeEvent === "roundStart"){

      } else if (obj.TypeEvent === "info"){

      } else if (obj.TypeEvent === "meteo"){

      } else if (obj.TypeEvent === "action"){

      } else if (obj.TypeEvent === "death"){

      } else if (obj.TypeEvent === "roundEnd"){

      } else if (obj.TypeEvent === "constructRaft"){

      } else if (obj.TypeEvent === "gameEnd"){
        this.setState({ gameEnd : true})
        client.close()
      }   
      client.send(JSON.stringify({
        fromPage : "simulation",
        message : "je suis toujours connecté :p",
      }));  
    }
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
    if(!this.state.gameStart){
      return (<div className="home">
      <PageDescription url_img={process.env.PUBLIC_URL + "/loading.gif"} page_title={"veuillez attendre"} descr_text={"........"} /></div>)
    }



    return (
      <div className="home">
        
        <div className="scrollmenu">
        <div className="players">
          <label>Joueurs restants :</label>
          <ul className="no-bullets">
              {this.state.players.map((player) => (
                <div className="flex-container">
                <li>
                    <p>{player.name}</p>
                    <PlayerImage role={player.role}/>
                    <p>Egoïsme : {player.selfishness}</p>
                    <p>Intelligence :  {player.intelligence}</p>
                </li>
                </div>
               ))}
          </ul>
        </div>
        </div>

        <div className="gameInfo">
          <div className="log">
            <label>Log : </label>
            <div className="gameLog">
            {this.state.messages.map((message) => (
              <p>{message}</p>
            ))}
            </div>
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
              <th>Place Radeau</th>
              <th>Stock Bois</th>
              <th>Stock Eau</th>
              <th>Stock StockNourriture</th>
            </tr>
            <tr>
              <td>{this.state.PlaceRadeau}</td>
              <td>{this.state.stockBois}</td>
              <td>{this.state.stockEau}</td>
              <td>{this.state.StockNourriture}</td>
            </tr>
            </table> 
          </div>
        </div>
       
    </div>
    );
  }

   renderPlayerImg(props) {
    if (props.role === ROLE_FISHERMAN) {    
      return <img src={process.env.PUBLIC_URL + "fishman.png"} width={"120px"} height={"200px"} alt="pêcheur"></img>;  
    } else if (props.role === ROLE_HANDYMAN) {    
      return <img src={process.env.PUBLIC_URL + "woodmaker.png"} width={"130px"} height={"200px"} alt="Météo"></img>;  
    } 
    return <img src={process.env.PUBLIC_URL + "normal_person.png"} width={"130px"} height={"200px"} alt="Météo"></img>;
  }
}


export default Game;