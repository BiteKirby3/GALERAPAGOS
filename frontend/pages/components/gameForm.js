import React from "react";


class GameForm extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            players : this.props.players
        };
    }

    render() {    
        console.log(this.state.players)        
        return (
            <div>
                {this.state.players.map((item)=>(<p key={item.name}>{item.name}</p>))}
            </div>
        );
    }
}

export default GameForm;