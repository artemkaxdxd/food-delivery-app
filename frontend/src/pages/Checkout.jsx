import React, { useEffect, useState } from 'react';
import { useSelector } from 'react-redux';
import { Container, Row, Col } from 'reactstrap';
import CommonSection from '../components/UI/common-section/CommonSection';
import Helmet from '../components/Helmet/Helmet';
import { useDispatch } from 'react-redux';

import { cartActions } from '../store/shopping-cart/cartSlice';
import '../styles/checkout.css';
import axios from 'axios';

const Checkout = () => {
  const [enterComment, setEnterComment] = useState('');
  const [enterEmail, setEnterEmail] = useState('');
  const [enterNumber, setEnterNumber] = useState('');
  const [enterAddress, setEnterAddress] = useState('');

  const shippingCost = 30;
  const dispatch = useDispatch();

  const cartTotalAmount = useSelector((state) => state.cart.totalAmount);
  const cartItems = useSelector((state) => state.cart.cartItems) || [];
  const totalAmount = cartTotalAmount + Number(shippingCost);

  const [userId, setUserId] = useState();
  const [userToken, setUserToken] = useState();
  useEffect(() => {
    const user = JSON.parse(localStorage.getItem('user'));
    setEnterEmail(user.email);
    setUserId(user.id);
    setUserToken(user.token);
  }, []);

  const [checkoutComplele, setCheckoutComplete] = useState(false);

  const submitHandler = async (e) => {
    e.preventDefault();

    try {
      const req = await axios.post(
        'http://localhost:8080/users/' + userId + '/orders',
        {
          address: enterAddress,
          phone_number: enterNumber,
          email: enterEmail,
          comment: enterComment,
          items: cartItems.map((item) => ({
            amount: item.quantity,
            item_id: item.id,
          })),
        },
        {
          headers: {
            Authorization: 'Bearer ' + userToken,
          },
        }
      );

      cartItems.forEach(({ id }) => {
        dispatch(cartActions.deleteItem(id));
      });

      setCheckoutComplete(true);
    } catch (error) {
      console.log(error);
    }
  };
  if (checkoutComplele) {
    return (
      <Helmet title="Checkout Complete">
        <CommonSection title="Checkout Complete" />
        <section>
          <Container>
            <h6>Order in process </h6>Once order is ready we will let you know
          </Container>
        </section>
      </Helmet>
    );
  }
  return (
    <Helmet title="Checkout">
      <CommonSection title="Checkout" />
      <section>
        <Container>
          <Row>
            <Col lg="8" md="6">
              <h6 className="mb-4">Shipping Address</h6>
              <form className="checkout__form" onSubmit={submitHandler}>
                <div className="form__group">
                  <input
                    type="email"
                    placeholder="Enter your email"
                    required
                    onChange={(e) => setEnterEmail(e.target.value)}
                    value={enterEmail}
                  />
                </div>
                <div className="form__group">
                  <input
                    type="number"
                    placeholder="Phone number"
                    required
                    onChange={(e) => setEnterNumber(e.target.value)}
                  />
                </div>
                <div className="form__group">
                  <input
                    type="text"
                    placeholder="Address"
                    required
                    onChange={(e) => setEnterAddress(e.target.value)}
                  />
                </div>
                <div className="form__group">
                  <input
                    type="text"
                    placeholder="Comment"
                    required
                    onChange={(e) => setEnterComment(e.target.value)}
                  />
                </div>
                <button type="submit" className="addTOCart__btn">
                  Payment
                </button>
              </form>
            </Col>

            <Col lg="4" md="6">
              <div className="checkout__bill">
                <h6 className="d-flex align-items-center justify-content-between mb-3">
                  Subtotal: <span>${cartTotalAmount}</span>
                </h6>
                <h6 className="d-flex align-items-center justify-content-between mb-3">
                  Shipping: <span>${shippingCost}</span>
                </h6>
                <div className="checkout__total">
                  <h5 className="d-flex align-items-center justify-content-between">
                    Total: <span>${totalAmount}</span>
                  </h5>
                </div>
              </div>
            </Col>
          </Row>
        </Container>
      </section>
    </Helmet>
  );
};

export default Checkout;
