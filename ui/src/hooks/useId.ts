let incrementalId = 0

export default function useId() {
    return incrementalId++
}