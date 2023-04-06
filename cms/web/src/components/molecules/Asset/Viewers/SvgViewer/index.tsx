import styled from "@emotion/styled";
import { useCallback, useEffect, useState } from "react";

import { useT } from "@reearth-cms/i18n";

type Props = { url: string; svgRender: boolean };

const SvgViewer: React.FC<Props> = ({ url, svgRender }) => {
  const t = useT();
  const [svgText, setSvgText] = useState("");

  const fetchData = useCallback(async () => {
    const res = await fetch(url, {
      method: "GET",
      headers: {
        "Access-Control-Allow-Origin": "*",
      },
    });
    if (res.status !== 200) {
      setSvgText(t("Could not display svg"));
      return;
    }
    const text = await res.text();
    setSvgText(text);
  }, [url, t]);

  useEffect(() => {
    fetchData();
  }, [fetchData]);

  return svgRender ? <Image src={url} alt="svg-preview" /> : <p>{svgText}</p>;
};

const Image = styled.img`
  width: 100%;
  height: 500px;
  object-fit: contain;
`;

export default SvgViewer;
