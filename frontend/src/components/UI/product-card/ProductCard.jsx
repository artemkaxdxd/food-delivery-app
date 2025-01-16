import React, { useEffect, useState } from 'react';

import '../../../styles/product-card.css';

import { Link } from 'react-router-dom';

import { useDispatch } from 'react-redux';
import { cartActions } from '../../../store/shopping-cart/cartSlice';

const ProductCard = (props) => {
  const { id, title, image_url, price } = props.item;
  const dispatch = useDispatch();

  const [image, setImage] = useState();

  useEffect(() => {
    setImage(image_url);
  }, [image_url]);

  const addToCart = () => {
    dispatch(
      cartActions.addItem({
        id,
        title,
        image_url,
        price,
      })
    );
  };

  return (
    <div className="product__item">
      <div className="product__img">
        <img src={image} alt="product-img" className="w-50" />
      </div>

      <div className="product__content">
        <h5>
          <Link to={`/foods/${id}`}>{title}</Link>
        </h5>
        <div className=" d-flex align-items-center justify-content-between ">
          <span className="product__price">${price}</span>
          <button className="addTOCart__btn" onClick={addToCart}>
            Add to Cart
          </button>
        </div>
      </div>
    </div>
  );
};

export default ProductCard;
