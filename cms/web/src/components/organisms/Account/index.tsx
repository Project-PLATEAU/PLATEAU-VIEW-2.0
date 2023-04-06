import Loading from "@reearth-cms/components/atoms/Loading";
import AccountSettingsMolecule from "@reearth-cms/components/molecules/AccountSettings";

import useHooks from "./hooks";

const AccountSettings: React.FC = () => {
  const { me, loading, handleUserUpdate, handleLanguageUpdate, handleUserDelete } = useHooks();

  return !me || loading ? (
    <Loading minHeight="400px" />
  ) : (
    <AccountSettingsMolecule
      user={me}
      onUserUpdate={handleUserUpdate}
      onLanguageUpdate={handleLanguageUpdate}
      onUserDelete={handleUserDelete}
    />
  );
};

export default AccountSettings;
