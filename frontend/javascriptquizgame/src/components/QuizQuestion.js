import React, { Component } from 'react'
import Highlight from 'react-highlight.js'

import './css/QuizQuestion.css'

export class QuizQuestion extends Component {

    submitAnswer = (answer) => {
        if (answer === this.props.question.correct_answer) {
            document.getElementById('correctAnswerAlert').style.display = 'block'
            this.props.incrementScore()
        } else {
            document.getElementById('wrongAnswerAlert').style.display = 'block'
        }
        
        setTimeout(() => {
            document.getElementById('correctAnswerAlert').style.display = 'none'
            document.getElementById('wrongAnswerAlert').style.display = 'none'
            this.props.setNextQuestion()
        }, 1000);
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

    render(){
        if(this.props.question !== undefined) {
            console.log(this.props.question)

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
                </div>
            )
        } else { // wait till the data comes in
            return(<div></div>)
        }

    }
}

export default QuizQuestion