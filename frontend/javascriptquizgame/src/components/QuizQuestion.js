import React, { Component } from 'react'
import Highlight from 'react-highlight.js'
import  { Markdown } from 'react-showdown'

import './css/QuizQuestion.css'

export class QuizQuestion extends Component {

    alreadySubmitted = 0

    submitAnswer = (answer) => {
        if (answer === this.props.question.correct_answer) {
            // Display the alert for correct  & increment score
            document.getElementById('correctAnswerAlert').style.display = 'block'
            
            if(this.alreadySubmitted === 0){
                this.props.incrementScore()
            }
        } else {
            // Display the alert for wrong answer
            document.getElementById('wrongAnswerAlert').style.display = 'block'
        }

        this.alreadySubmitted = 1
        
        this.displayExplanation()

        setTimeout(() => { // Hide the prompts
            document.getElementById('correctAnswerAlert').style.display = 'none'
            document.getElementById('wrongAnswerAlert').style.display = 'none'
        }, 2000);
    }

    fillOptions = (options) => {
        let optionsHTML = []

        const buttonStyle = {
            marginBottom: '0rem', 
            marginTop: '1rem',
            width: '100%'
        }

        for(let key in options) {
            // To remove the ` on the ends
            let optionText = options[key].replace(/`/g , "")

            optionsHTML.push(
                <button onClick={() => this.submitAnswer(key.toString())} className="btn btn-primary" style={buttonStyle} key={key}>
                    {key}. {optionText}
                </button>
                )
            optionsHTML.push(<br key={key.toString() + '_br'}></br>)
        }

        return optionsHTML
    }

    displayExplanation = () => {
        document.getElementById("questionExplanation").style.display = 'block';
    }

    hideExplanation = () => {
        document.getElementById("questionExplanation").style.display = 'none';
    }


    render(){
        if(this.props.question !== undefined) {
            
            return (
                <div>
                    <div id="correctAnswerAlert" style={{display: 'none'}}>
                        <div className="alert alert-info alert-dismissible fade show" role="alert">
                            <strong>That Correct!</strong>
                        </div>
                    </div>

                    <div id="wrongAnswerAlert" style={{display: 'none'}}>
                        <div className="alert alert-danger alert-dismissible fade show" role="alert">
                            <strong>That's Wrong!</strong>
                        </div>
                    </div>

                    <div className="questionCard">
                        <br></br>
                        <h3 className="statement">
                            <span className="number">{this.props.question.number}.</span> <br></br> {this.props.question.statement}
                        </h3>
                        <div className="code">
                            <Highlight language="javascript">{this.props.question.code}</Highlight>
                        </div>
                        { this.fillOptions(this.props.question.options) }
                    </div>

                    <div className="questionExplanation" id="questionExplanation">
                        <br></br>
                        <Markdown markup={ this.props.question.explanation }></Markdown>
                        <br></br>
                        <button onClick={ () => {this.hideExplanation(); this.props.setNextQuestion()}} className="btn btn-sm btn-warning">
                            Next Question
                        </button>

                    </div>
                </div>
            )
        } else { // wait till the data comes in
            return(<div></div>)
        }

    }
}

export default QuizQuestion