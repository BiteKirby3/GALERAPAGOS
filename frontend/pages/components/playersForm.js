import React from "react";
import styles from '../../styles/Players.module.css';
class PlayersForm extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            name : "",
            selfishness : 0, 
            intelligence : 0,     
        };
    this.handleChangeName = this.handleChangeName.bind(this);
    this.handleChangeSelfishness = this.handleChangeSelfishness.bind(this);
    this.handleChangeIntelligence = this.handleChangeIntelligence.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
    }


    handleSubmit(event) {
        event.preventDefault();
      }

    handleChangeName(event){
        this.setState({Name: event.target.value});
    }
    handleChangeSelfishness(event){
        this.setState({selfishness: event.target.value});
    }
    handleChangeIntelligence(event){
        this.setState({intelligence: event.target.value});
    }
  
    render() {
      return (
        <form>
            <div className={styles.scrollmenu}>
                <div>
                    Players 1 :
                    <p>
                    <label>
                        Nom : <input type="text" value={this.state.name} onChange={this.handleChangeName} />     
                    </label>
                    </p>
                    <p>
                    <label>
                        Rôle : <select>
                            <option selected value="none">Rien</option>
                            <option value="fisherman">Pêcheur</option>
                            <option value="handyman">Bricoloeur</option>
                        </select>    
                    </label>
                    </p>
                    <p>
                        <label>
                            Egoïsme : <input type="number" value={this.state.selfishness}  min={0} max={10} step={1} onChange={this.handleChangeSelfishness}/>   
                        </label>
                    </p>
                    <p>
                        <label>
                            Intelligence : <input type="number" value={this.state.intelligence}  min={0} max={10} step={1} onChange={this.handleChangeIntelligence}/>   
                        </label>
                    </p>
                </div>
            </div>
          <input type="submit" value="Suivant" />
        </form>
      );
    }
}

  export default PlayersForm;