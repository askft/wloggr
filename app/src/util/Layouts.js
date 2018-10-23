import React from "react";
import Logo from "../components/Logo";
import PublicNav from "../components/PublicNav";
import PrivateNav from "../components/PrivateNav";

export const PublicLayout = Component => props => (
  <div>
    <Logo />
    <PublicNav />
    <div className="fade-in">
      <Component {...props} />
    </div>
  </div>
);

export const PrivateLayout = Component => props => (
  <div>
    <Logo />
    <PrivateNav />
    <div className="fade-in">
      <Component {...props} />
    </div>
  </div>
);
