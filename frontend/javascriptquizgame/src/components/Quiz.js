import React, { Component } from 'react'

import './css/Quiz.css'

export class Quiz extends Component {

    state = {
        data: {}
    }

    componentDidMount = () => {
        //console.log("Getting questions")
        fetch("http://localhost:1200/api/getQuestions")
            .then(res => res.json())
            .then((json) => {
                this.setState({data: json})
                //console.log(json)
            })
    }

    
    render() {
        return (
            <div>
                <h1>I am quiz</h1>
            </div>
        )
    }
}

export default Quiz
