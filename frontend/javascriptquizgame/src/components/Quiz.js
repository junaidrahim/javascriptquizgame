import React, { Component } from 'react'
import QuizQuestion from './QuizQuestion'

import './css/Quiz.css'

export class Quiz extends Component {

    state = {
        data: {},
        currentQuestion: 0,
        endQuiz : 0
    }

    totalScore = 0;

    incrementScore = () => {
        this.totalScore += 1
        console.log(this.totalScore)
    }

    setNextQuestion = () => {
        if(this.state.currentQuestion+1 < this.state.data.length){
            this.setState({
                currentQuestion : this.state.currentQuestion + 1
            })
        } else {
            this.endQuiz()
        }
    }

    endQuiz = () => {
        console.log("end quiz ran")
        this.setState({
            endQuiz : 1
        })
    }

    componentDidMount = () => {
        //console.log("Getting questions")
        fetch("http://192.168.43.44:1200/api/getQuestions")
            .then(res => res.json())
            .then((json) => {
                this.setState({data: json})
                //console.log(json)
            })
            .catch(err => {
                alert("Server Unavailable, Please try again later")
            })
    }

    
    render() {
        if(this.state.endQuiz){
            return(
                <div>
                    <div className="scoreContainer">
                        <h1>Your total score is {this.totalScore}</h1>
                        <div style={{textAlign: 'center'}}>
                            <button className="btn btn-danger" onClick={() => document.location.reload()}>Play Again</button>
                        </div>
                    </div>
                </div>
            )

        } else {
            
            return (
                <div>
                    <div className="questionContainer">
                        <QuizQuestion 
                            question={this.state.data[this.state.currentQuestion]} 
                            incrementScore={this.incrementScore} 
                            setNextQuestion={this.setNextQuestion}>
                        </QuizQuestion>

                        <div style={{textAlign: 'center'}}>
                            <br></br>
                            <button className="btn btn-dark btn-sm" onClick={this.endQuiz}>End Game</button>
                            <br></br>
                            <br></br>
                        </div>
                    </div>

                </div>
            )
        
        }
    }
}

export default Quiz
