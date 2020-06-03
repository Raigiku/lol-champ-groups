import React, { useState } from 'react'
import { FormControl, InputLabel, Select, MenuItem, FormHelperText, Typography } from '@material-ui/core'
import { SelectItem } from '../../entities/common/SelectItem'

type Props = {
  required: boolean
  items: SelectItem[]
  label: string
  name: string
  helperText: string
  updateState: (prop: string, value?: string) => void
}

const SelectWrapper = ({
  required,
  items,
  label,
  name,
  helperText,
  updateState,
}: Props) => {
  const [inputValue, setInputValue] = useState('')
  const [error, setError] = useState<string | undefined>('Required')

  const handleChange = (e: React.ChangeEvent<{ name?: string | undefined; value: unknown; }>) => {
    const value = e.target.value as string
    const name = e.target.name as string

    setInputValue(value)
    setError(undefined)
    updateState(name, value)
  }

  return (
    <FormControl required={required} error={error != null}>
      <InputLabel>{label}</InputLabel>
      <Select
        value={inputValue}
        onChange={handleChange}
        name={name}
      >
        {items.map(item =>
          <MenuItem key={item.value} value={item.value}>{item.label}</MenuItem>
        )}
      </Select>
      <FormHelperText>
        <>
          {error &&
            <>
              <Typography variant='caption'>{error}</Typography>
              <br />
            </>
          }
          <Typography variant='caption'>{helperText}</Typography>
        </>
      </FormHelperText>
    </FormControl>
  )
}

export default SelectWrapper
