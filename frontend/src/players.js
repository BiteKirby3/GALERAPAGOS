import React from "react";
import PageDescription from "./pageDescription";
import { NavLink } from "react-router-dom";
import styles from './players.css';

const ROLE_FISHERMAN = "fisherman"
const ROLE_HANDYMAN = "handyman"
const ROLE_NONE = "none"

class Players extends React.Component {
  constructor(props) {
    super(props);
    //console.log(this.props.location);
    const players = []
    for (var i = 0; i < this.props.nbPlayers; i++) {
        players.push({
            id : i,
            name : "",
            role : ROLE_NONE,
            selfishness : 0, 
            intelligence : 0,   
        });
    }
    this.state = {
        players  
    };
    this.handleChangeName = this.handleChangeName.bind(this);
    this.handleChangeSelfishness = this.handleChangeSelfishness.bind(this);
    this.handleChangeIntelligence = this.handleChangeIntelligence.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
}


handleSubmit(event) {
    event.preventDefault();
    this.props.onSubmit(this.state.players);
  }

updatePlayer(index, attributes){
    let players = [...this.state.players];
    let player = {
        ...players[index],
        ...attributes
    }
    players[index] = player;
    this.setState({players});
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
    console.log(this.props);
    return (
      <div className="home">
        <PageDescription url_img={process.env.PUBLIC_URL + "/actions.jpg"} page_title={"     TO ADD : ................."} descr_text={"    Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod\n tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, \nquis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. \nDuis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore \neu fugiat nulla pariatur. Excepteur sint occaecat cupidatat \nnon proident, sunt in culpa qui officia deserunt mollit anim id est\n laborum."} />
        <div className={styles.scrollmenu}>
                {Array.from({length: this.props.nbPlayers},(_, i) => (
                    <div key={i}>
                        Players {i+1} :
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
                                <option value={ROLE_HANDYMAN}>Bricoloeur</option>
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
                ))     
                }
            </div>
          <NavLink className="nav-link" to="/simulation">
              <input type="submit" value="Lancer Simulation" />
          </NavLink>   
    </div>
    );
  }
}

export default Players;