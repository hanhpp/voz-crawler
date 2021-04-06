import { Button } from '@material-ui/core';
import axios from 'axios';
import * as React from "react";

class Body extends React.Component {
    constructor(props) {
        super(props);
        this.state = {};
        this.handleClick.bind(this.handleClick())
    }

    handleClick() {
        var param = 'threads'
        const axios = require('axios').default;

        var url = "http://localhost:9500/"
        axios.get(url + param)
            .then(res => function () {
                var data = res.data[0]
                alert(data)
                console.log(data)
            }())
            .catch(
                (error) => {
                    console.log(error)
                })
            .then(console.log("Finished!"))
    }

    render () {
        return (
            <main className={this.props.main}>
                <h1 className={this.props.title}>
                    Welcome Voz reader
                </h1>
                <Button onClick={this.handleClick} color="primary">Click me</Button>
            </main>
        )
    }
}

export default Body