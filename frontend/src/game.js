import React from "react";
import { Link } from "react-router-dom";
import { w3cwebsocket as W3CWebSocket } from "websocket";

const client = new W3CWebSocket('ws://127.0.0.1:5000');


class Game extends React.Component {
  constructor(props) {
    super(props);
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

  render() {
    return (
        <div className="home">
          <div className="container">
            <Link to="/simulation/this-is-a-player">
              <div className="row align-items-center my-5">
                <div className="col-lg-7">
                  <img
                      className="img-fluid rounded mb-4 mb-lg-0"
                      src="http://placehold.it/900x400"
                      alt=""
                  />
                </div>
                <div className="col-lg-5">
                  <h1 className="font-weight-light">This is a post title</h1>
                  <p>
                    Lorem Ipsum is simply dummy text of the printing and typesetting
                    industry. Lorem Ipsum has been the industry's standard dummy
                    text ever since the 1500s, when an unknown printer took a galley
                    of type and scrambled it to make a type specimen book.
                  </p>
                </div>
              </div>
            </Link>
          </div>
        </div>
    );
  }
}
export default Game;