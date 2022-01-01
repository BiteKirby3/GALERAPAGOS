import React from "react";
import './navigation.css';
import { NavLink } from "react-router-dom";


//used to disabled the NavLink in the NavBar
function handleClick (e) {
  e.preventDefault()
}

function Navigation() {

  return (
    <div className="navigation">
      <nav className="navbar navbar-expand navbar-dark">
        <div className="container">
          <div>
            <NavLink className="navbar-brand" to="/">
            <p className="title">A21-IA04-Projet : Simulation d'un jeu de société — Galèrapagos </p>
            <p className="slogan">Coopérer pour lutter mais trahir pour gagner ! </p>
            </NavLink>
          </div>
          <div>
            <ul className="navbar-nav ml-auto">
              <li className="nav-item">
                <NavLink className="nav-link" to="/" onClick={handleClick}>
                Galèrapagos
                  <span className="sr-only">(current)</span>
                </NavLink>
              </li>
              <li className="nav-item">
                <NavLink className="nav-link" to="/players" onClick={handleClick}>
                Joueurs
                </NavLink>
              </li>
              <li className="nav-item">
                <NavLink className="nav-link" to="/simulation" onClick={handleClick}>
                Simulation
                </NavLink>
              </li>
            </ul>
          </div>
        </div>
      </nav>
    </div>
  );
}


export default Navigation;