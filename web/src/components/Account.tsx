import React from 'react';
import { Box, Typography, List, ListItem, ListItemAvatar, Avatar, ListItemText, IconButton, Paper, Divider, ListItemIcon, Collapse, ListItemButton } from '@mui/material';
import DeleteIcon from '@mui/icons-material/Delete';
import FolderIcon from '@mui/icons-material/Folder';
import { styled } from '@mui/system';
import { IAccount } from '../models/types';
import AttachMoneySharpIcon from '@mui/icons-material/AttachMoneySharp';
import CurrencyRubleSharpIcon from '@mui/icons-material/CurrencyRubleSharp';
import MoreVertIcon from '@mui/icons-material/MoreVert';
import { ExpandLess, ExpandMore, StarBorder } from '@mui/icons-material';
import InboxIcon from '@mui/icons-material/MoveToInbox';
import ModeSharpIcon from '@mui/icons-material/ModeSharp';
import DeleteForeverSharpIcon from '@mui/icons-material/DeleteForeverSharp';

// Определение пропсов
interface AccountProps {
    title: string;
    accounts: IAccount[];
}

const Account: React.FC<AccountProps> = ({ title, accounts }) => {
    const [open, setOpen] = React.useState(
        Array(accounts.length).fill(false)
    );

    const handleClick = (index: number) => {
        // Копируем текущий массив состояний
        const newOpen = [...open];
        // Изменяем состояние для конкретного элемента
        newOpen[index] = !newOpen[index];
        // Обновляем состояние
        setOpen(newOpen);
    };

    const sortedAccounts = [...accounts].sort((a, b) => a.currency.localeCompare(b.currency));

    return (
        <Box sx={{ width: '100%' }}>
            <Typography variant="h6" component="div" sx={{ textAlign: 'center' }}>
                {title}
            </Typography>
            <div>
                <List>
                    {sortedAccounts.map((item, index) => (
                        <div>
                            <ListItem sx={{ boxShadow: '0px 3px 5px rgba(0, 0, 0, 0.2)', mb: 1 }} key={index}
                                secondaryAction={
                                    <ListItemButton onClick={() => handleClick(index)}>
                                        {open[index] ? <ExpandLess /> : <ExpandMore />}
                                    </ListItemButton>
                                }
                            >
                                <ListItemAvatar>
                                    <Avatar>
                                        {
                                            item.currency === '810' ? (<CurrencyRubleSharpIcon />) :
                                                item.currency === '840' ? (<AttachMoneySharpIcon />) : null
                                        }
                                    </Avatar>
                                </ListItemAvatar>
                                <ListItemText
                                    primary={item.name}
                                    secondary={item.amount.toLocaleString('ru-Ru', { minimumFractionDigits: 2, maximumFractionDigits: 2 })}
                                />

                            </ListItem>
                            <Collapse in={open[index]} timeout="auto" unmountOnExit>
                                <List component="div" disablePadding>
                                    <ListItemButton sx={{ display: 'flex', justifyContent: 'space-around' }}>
                                        <ListItemText sx={{ ml: 7 }} primary={item.account} secondary={item.interestRate * 100 + '%'} />
                                        <Box sx={{ display: 'flex', flexDirection: 'column', alignItems: 'center', mr: 1 }}>
                                            <IconButton size='small'>
                                                <ModeSharpIcon />
                                            </IconButton>
                                            <IconButton size='small'>
                                                <DeleteIcon />
                                            </IconButton>
                                        </Box>
                                    </ListItemButton>
                                </List>
                            </Collapse>
                        </div>


                    ))}
                </List>
            </div>
        </Box>
    );
};

export default Account;
