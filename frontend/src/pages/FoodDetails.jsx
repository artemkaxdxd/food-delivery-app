import React, { useState, useEffect } from 'react';

import { useParams } from 'react-router-dom';
import Helmet from '../components/Helmet/Helmet';
import CommonSection from '../components/UI/common-section/CommonSection';
import { Container, Row, Col } from 'reactstrap';

import { useDispatch } from 'react-redux';
import { cartActions } from '../store/shopping-cart/cartSlice';

import '../styles/product-details.css';

const FoodDetails = () => {
  const [loaded, setLoaded] = useState(false);

  const [tab, setTab] = useState('desc');
  const [enteredName, setEnteredName] = useState('');
  const [enteredEmail, setEnteredEmail] = useState('');
  const [reviewMsg, setReviewMsg] = useState('');
  const { id } = useParams();
  const dispatch = useDispatch();

  const [item, setItem] = useState();
  useEffect(() => {
    let localItems = JSON.parse(localStorage.getItem('items'));
    const item = localItems.filter((item) => item.id === parseInt(id))[0];
    setItem(item);
    setLoaded(true);
  }, [id, setLoaded, dispatch]);

  const addToCart = () => {
    dispatch(
      cartActions.addItem({
        id: item.id,
        title: item.title,
        image_url: item.image_url,
        price: item.price,
      })
    );
  };

  const submitHandler = (e) => {
    e.preventDefault();

    console.log(enteredName, enteredEmail, reviewMsg);
  };

  useEffect(() => {
    window.scrollTo(0, 0);
  }, [item]);

  if (!loaded) return <Helmet title="Product-details"> </Helmet>;
  return (
    <Helmet title="Product-details">
      <CommonSection title={item.title} />

      <section>
        <Container>
          <Row>
            <Col lg="2" md="2">
              <div className="product__images ">
                <div className="img__item mb-3">
                  <img src={item.image_url} alt="" className="w-50" />
                </div>
              </div>
            </Col>

            <Col lg="4" md="4">
              <div className="product__main-img">
                <img src={item.image_url} alt="" className="w-100" />
              </div>
            </Col>

            <Col lg="6" md="6">
              <div className="single__product-content">
                <h2 className="product__title mb-3">{item.title}</h2>
                <p className="product__price">
                  {' '}
                  Price: <span>${item.price}</span>
                </p>
                <button onClick={addToCart} className="addTOCart__btn">
                  Add to Cart
                </button>
              </div>
            </Col>

            <Col lg="12">
              <div className="tabs d-flex align-items-center gap-5 py-3">
                <h6
                  className={` ${tab === 'desc' ? 'tab__active' : ''}`}
                  onClick={() => setTab('desc')}
                >
                  Description
                </h6>
                <h6
                  className={` ${tab === 'rev' ? 'tab__active' : ''}`}
                  onClick={() => setTab('rev')}
                >
                  Review
                </h6>
              </div>

              {tab === 'desc' ? (
                <div className="tab__content">
                  <p>{item.desciption}</p>
                </div>
              ) : (
                <div className="tab__form mb-3">
                  <div className="review pt-5">
                    <p className="user__name mb-0">Jhon Doe</p>
                    <p className="user__email">jhon1@gmail.com</p>
                    <p className="feedback__text">great product</p>
                  </div>

                  <div className="review">
                    <p className="user__name mb-0">Jhon Doe</p>
                    <p className="user__email">jhon1@gmail.com</p>
                    <p className="feedback__text">great product</p>
                  </div>

                  <div className="review">
                    <p className="user__name mb-0">Jhon Doe</p>
                    <p className="user__email">jhon1@gmail.com</p>
                    <p className="feedback__text">great product</p>
                  </div>
                  <form className="form" onSubmit={submitHandler}>
                    <div className="form__group">
                      <input
                        type="text"
                        placeholder="Enter your name"
                        onChange={(e) => setEnteredName(e.target.value)}
                        required
                      />
                    </div>

                    <div className="form__group">
                      <input
                        type="text"
                        placeholder="Enter your email"
                        onChange={(e) => setEnteredEmail(e.target.value)}
                        required
                      />
                    </div>

                    <div className="form__group">
                      <textarea
                        rows={5}
                        type="text"
                        placeholder="Write your review"
                        onChange={(e) => setReviewMsg(e.target.value)}
                        required
                      />
                    </div>

                    <button type="submit" className="addTOCart__btn">
                      Submit
                    </button>
                  </form>
                </div>
              )}
            </Col>
          </Row>
        </Container>
      </section>
    </Helmet>
  );
};

export default FoodDetails;
