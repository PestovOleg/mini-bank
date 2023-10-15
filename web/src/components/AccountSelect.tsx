import * as React from 'react';
import Box from '@mui/material/Box';
import TextField from '@mui/material/TextField';
import Autocomplete from '@mui/material/Autocomplete';
import { IAccount } from '../models/types';
import { Avatar } from '@mui/material';

interface AccountSelectProps {
  accounts: IAccount[];
  placeHolder: string;
  style?: React.CSSProperties;
  onAccountSelected: (selectedAccount: IAccount | null) => void;
}

const AccountSelect: React.FC<AccountSelectProps> = ({ accounts, placeHolder,onAccountSelected  }) => {
  return (
    <Autocomplete
      id="account-select-demo"
      sx={{ marginBottom: 1, marginTop: 1 }}
      fullWidth
      options={accounts}
      isOptionEqualToValue={(option, value) => option.id === value.id}
      onChange={(event, newValue) => {
        onAccountSelected(newValue);
      }}
      autoHighlight
      getOptionLabel={(option) => option.account}
      renderOption={(props, option) => (
        <Box component="li" sx={{ display: 'flex', alignItems: 'center' }} {...props}>
          <Avatar sx={{ width: 36, height: 36, marginRight: 2 }} >
            {
              option.currency === '810' ? (<img alt="Ruble" src="/ruble.png" />) :
                option.currency === '840' ? (<img alt="Dollar" src="/dollar.png" />) : null
            }
          </Avatar>
          <Box sx={{ flexGrow: 1, display: 'flex', flexDirection: 'column', marginLeft: 2 }}>
            <Box sx={{ display: 'flex', justifyContent: 'space-between' }}>
              <span>{option.name}</span>
              <span>*{option.account.slice(-4)}</span>
            </Box>
            <Box sx={{ alignSelf: 'flex-start' }}>
              {option.amount.toLocaleString('ru-RU', { minimumFractionDigits: 2, maximumFractionDigits: 2 })}
            </Box>
          </Box>
        </Box>

      )}
      renderInput={(params) => (
        <TextField
          {...params}
          label={placeHolder}
          inputProps={{
            ...params.inputProps,
            autoComplete: 'new-password',
          }}
        />
      )}
    />
  );
}

export default AccountSelect;