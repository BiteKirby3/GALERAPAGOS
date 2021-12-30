import React from "react";


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
            <div class="description">
                    <img
                        class="page_image"
                        src={this.state.url_img}
                        alt=""
                    />
            <pre>
                <h1 class="font-weight-light">{this.state.page_title}</h1>
                {this.state.descr_text}
            </pre>
        </div>
        );
    }
}

export default PageDescription;