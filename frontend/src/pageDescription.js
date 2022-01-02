import React from "react";
import './pageDescription.css';


class PageDescription extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            page_title : this.props.page_title,
            url_img : this.props.url_img,
            descr_text : this.props.descr_text,
        };
    }

    render() {         
        return (
            <div className="description">
                <img
                    className="page_image"
                    src={this.state.url_img}
                    alt=""
                />
                <div>
                    <h1 className="font-weight-light">{this.state.page_title}</h1>
                        <h5>{this.state.descr_text}</h5>
                </div>
            </div>
        );
    }
}

export default PageDescription;