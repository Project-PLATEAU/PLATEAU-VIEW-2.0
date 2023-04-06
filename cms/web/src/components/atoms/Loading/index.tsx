import Row from "@reearth-cms/components/atoms/Row";
import Spin from "@reearth-cms/components/atoms/Spin";
import { useT } from "@reearth-cms/i18n";

export type Props = {
  spinnerSize?: "small" | "large" | "default";
  minHeight?: string;
};

const Loading: React.FC<Props> = ({ spinnerSize, minHeight }) => {
  const t = useT();

  return (
    <Row justify="center" align="middle" style={{ minHeight }}>
      <Spin tip={t("Loading")} size={spinnerSize} />
    </Row>
  );
};

export default Loading;
