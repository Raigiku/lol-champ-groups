import React, { useState } from 'react'
import { FormControl, FormLabel, FormControlLabel, Checkbox, FormGroup, FormHelperText, Typography } from '@material-ui/core'
import { CheckboxGroupItem } from '../../entities/common/CheckboxGroupItem'

type Props = {
  items: CheckboxGroupItem[]
  required: boolean
  label: string
  name: string
  helperText: string
  updateState: (prop: string, value?: Set<string>) => void
}

const CheckboxGroupWrapper = ({
  items,
  required,
  helperText,
  label,
  name,
  updateState
}: Props) => {
  const [inputValues, setInputValues] = useState<Set<string>>(new Set())
  const [error, setError] = useState<string | undefined>('Required')

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const value = e.target.name
    setInputValues(prev => {
      if (prev.has(value)) {
        prev.delete(value)
      } else {
        prev.add(value)
      }

      const newError = (required && prev.size === 0) ? 'Required' : undefined
      setError(newError)

      const stateValue = newError != null ? undefined : prev
      updateState(name, stateValue)

      return new Set(prev)
    })
  }

  return (
    <FormControl required={required} error={error != null}>
      <FormLabel>{label}</FormLabel>
      <FormGroup>
        {items.map(item =>
          <FormControlLabel
            key={item.name}
            label={item.label}
            control={
              <Checkbox
                checked={inputValues.has(item.name)}
                onChange={handleChange}
                name={item.name}
              />
            }
          />
        )}
      </FormGroup>
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

export default CheckboxGroupWrapper
