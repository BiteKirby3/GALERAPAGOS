import React, { useEffect } from "react";

const ROLE_FISHERMAN = "fisherman"
const ROLE_HANDYMAN = "handyman"

function PlayerImage(props) {
  if (props.role === ROLE_FISHERMAN) {    
    return <img src={process.env.PUBLIC_URL + "fishman.png"} width={"120px"} height={"200px"} alt="pêcheur"></img>;  
  } else if (props.role === ROLE_HANDYMAN) {    
    return <img src={process.env.PUBLIC_URL + "woodmaker.png"} width={"130px"} height={"200px"} alt="Météo"></img>;  
  } 
  return <img src={process.env.PUBLIC_URL + "normal_person.png"} width={"130px"} height={"200px"} alt="Météo"></img>;
}


export default PlayerImage;