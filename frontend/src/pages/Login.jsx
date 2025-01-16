import React, { useRef } from 'react';
import Helmet from '../components/Helmet/Helmet';
import CommonSection from '../components/UI/common-section/CommonSection';
import { Container, Row, Col } from 'reactstrap';
import { Link, useNavigate } from 'react-router-dom';
import axios from 'axios';
import { jwtDecode } from 'jwt-decode';

const Login = () => {
  const loginNameRef = useRef();
  const loginPasswordRef = useRef();

  const navigator = useNavigate();
  const submitHandler = async (e) => {
    e.preventDefault();
    const [email, password] = [
      loginNameRef.current.value,
      loginPasswordRef.current.value,
    ];
    try {
      const req = await axios.post('http://localhost:8080/users/login', {
        email,
        password,
      });
      const token = req.data.data.token;
      const { id } = jwtDecode(token);

      const getMe = await axios.get('http://localhost:8080/users/' + id, {
        headers: { Authorization: `Bearer ` + token },
      });
      localStorage.setItem(
        'user',
        JSON.stringify({
          email,
          password,
          id,
          token,
          name: getMe.data.data.token.name,
        })
      );
      navigator('/foods');
    } catch (error) {
      console.log(error.response.data.description);
    }
  };

  return (
    <Helmet title="Login">
      <CommonSection title="Login" />
      <section>
        <Container>
          <Row>
            <Col lg="6" md="6" sm="12" className="m-auto text-center">
              <form className="form mb-5" onSubmit={submitHandler}>
                <div className="form__group">
                  <input
                    type="email"
                    placeholder="Email"
                    required
                    ref={loginNameRef}
                  />
                </div>
                <div className="form__group">
                  <input
                    type="password"
                    placeholder="Password"
                    required
                    ref={loginPasswordRef}
                  />
                </div>
                <button type="submit" className="addTOCart__btn">
                  Login
                </button>
              </form>
              <Link to="/register">
                Don't have an account? Create an account
              </Link>
            </Col>
          </Row>
        </Container>
      </section>
    </Helmet>
  );
};

export default Login;
