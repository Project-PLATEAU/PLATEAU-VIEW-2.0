import styledComponents, { StyledInterface } from "styled-components";

import common from "./commonStyles";

export const styled = import.meta.env.DEV ? styledComponents : (window.styled as StyledInterface);

export const commonStyles = common;
