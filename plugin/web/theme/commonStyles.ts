const common = {
  mainWrapper: `
    width: 100%;
    height: 100%;
    transition: height 0.5s, width 0.5s;
    border-radius: 0;
  `,
  minimizedWrapper: `
    width: 226px;
    height: 80px;
    border-radius: 0 0 8px 0;
  `,
  simpleButton: `
    display: flex;
    justify-content: center;
    align-items: center;
    background: #ffffff;
    border: 1px solid #d9d9d9;
    border-radius: 2px;
    padding: 5px;
    height: 24px;
    width: 51px;
    align-self: flex-end;
    cursor: pointer;

    :hover {
      background: #f4f4f4;
    }
  `,
};

export default common;
