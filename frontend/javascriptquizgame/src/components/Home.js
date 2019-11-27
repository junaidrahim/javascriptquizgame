import React, { Component } from 'react'
import Quiz from './Quiz'

import './css/Home.css'

export class Home extends Component {

    state = {
        showQuiz : false
    }

    displayQuiz = () => {
        this.setState({
            showQuiz: true
        })
    }

    render(){
        if(this.state.showQuiz){
            return (
                <div>
                    <Quiz></Quiz>
                </div>
            )
        } else {
            return (
                <div className="home">
                    <div className="title-header">
                        <h1>the <br></br> javascript <br></br> quiz<span className="dot">.</span></h1>
                    </div>
                    <div className="content">
                        <br></br>
                        <p>Just a gamified version of the
                        javascript questions by <a target="_blank" rel="noopener noreferrer" href="https://github.com/lydiahallie/javascript-questions">Lydia Halie</a> <br></br>
                        A project by Junaid Rahim. <br></br>
                        <a style={{fontSize: '0.8rem'}} target="_blank" rel="noopener noreferrer" href="https://github.com/junaidrahim/javascriptquizgame">Star this on GitHub</a>
                        </p>
                        <br></br>
                    </div>
                    
                    <div className="button-container">
                        <button className="btn btn-danger btn-lg" onClick={this.displayQuiz} >Start Playing</button>
                    </div>
                </div>
            )
        }
    }

}

export default Home