import React, { useState } from 'react'
import { TextField, Typography } from '@material-ui/core'

type Props = {
  required: boolean
  label: string
  type?: string
  name: string
  helperText: string
  updateState: (prop: string, value?: string) => void
  isValueOk: (value: string) => string | undefined
}

const TextFieldWrapper = ({
  required,
  label,
  type,
  name,
  updateState,
  helperText,
  isValueOk,
}: Props) => {
  const [inputValue, setInputValue] = useState('')
  const [error, setError] = useState<string | undefined>('Required')

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    const { name, value } = e.currentTarget
    setInputValue(value)

    const newError = isValueOk(value)
    setError(newError)

    const stateValue = newError != null ? undefined : value
    updateState(name, stateValue)
  }

  return (
    <TextField
      required={required}
      label={label}
      type={type}
      name={name}
      value={inputValue}
      onChange={handleChange}
      error={error != null}
      helperText={
        <>
          {error &&
            <>
              <Typography variant='caption'>{error}</Typography>
              <br />
            </>
          }
          <Typography variant='caption'>{helperText}</Typography>
        </>
      }
    />
  )
}

export default TextFieldWrapper
