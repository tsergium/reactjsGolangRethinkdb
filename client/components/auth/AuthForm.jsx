import React, {Component} from 'react';

class AuthForm extends Component{
    onSubmit(e){
        e.preventDefault();
        const userNode = this.refs.userName;
        const passNode = this.refs.passWord;
        const username = userNode.value;
        const password = passNode.value;
        this.props.setUsername(username);
        this.props.setPassword(password);
        node.value = '';
    }
    render(){
        return (
            <form onSubmit={this.onSubmit.bind(this)}>
                <div className='form-group'>
                    <input
                        ref='userName'
                        type='text'
                        className='form-control'
                        placeholder='Set Your Name...' />
                    <input
                        ref='passWord'
                        type='password'
                        className='form-control'
                        placeholder='Set Your Password...' />
                    <input
                        type='submit'
                        className='btn btn-success'
                        value='Login'/>
                </div>
            </form>
        )
    }
}

AuthForm.propTypes = {

};

export default AuthForm