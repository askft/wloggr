import React from "react";
import { Link } from "react-router-dom";

const Logo = () => (
  <h1>
    <Link to="/" style={{ textDecoration: "none" }}>
      wloggr 🏋️
    </Link>
  </h1>
);

export default Logo;
