import { StyleSheet } from 'react-native';

export const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: 'white',
    padding: 16,
  },
  subContainer: {
    flexDirection: 'column',
    marginBottom: 16,
    padding: 16,
    borderRadius: 10,
  },
  header: {
    fontSize: 24,
    fontWeight: 'bold',
    marginBottom: 16,
    textAlign: 'center',
  },
  detailContainer: {
    flexDirection: 'row',
    marginBottom: 8,
  },
  ActionReactionHeader: {
    fontSize: 20,
    fontWeight: 'bold',
    textAlign: 'center',
    alignItems: 'center',
    marginBottom: 16,
  },
  label: {
    fontWeight: 'bold',
    fontSize: 16,
    color: 'white',
    marginRight: 8,
  },
  value: {
    color: 'white',
    fontSize: 16,
  },
  modalContainer: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
    backgroundColor: 'rgba(0, 0, 0, 0.5)',
  },
  modalContent: {
    padding: 20,
    borderRadius: 10,
    width: '80%',
  },
  modalHeader: {
    fontSize: 20,
    fontWeight: 'bold',
    color: 'white',
    marginBottom: 16,
    textAlign: 'center',
  },
  input: {
    borderWidth: 1,
    borderColor: 'gray',
    borderRadius: 8,
    padding: 8,
    marginBottom: 16,
  },
  cancelButton: {
    borderWidth: 2,
    borderColor: '#E60000',
    backgroundColor: 'white',
    borderRadius: 8,
    padding: 4,
    alignItems: 'center',
    marginBottom: 8,
  },
  saveButton: {
    borderWidth: 2,
    borderColor: 'white',
    borderRadius: 8,
    padding: 4,
    alignItems: 'center',
  },
  optionRow: {
    flexDirection: 'row',
    alignItems: 'center',
    marginBottom: 15,
    width: '100%',
  },
  optionLabel: {
    color: 'white',
    fontSize: 16,
    flex: 1,
  },
  optionInput: {
    flex: 2,
    backgroundColor: '#f0f0f0',
    borderRadius: 5,
    padding: 10,
    borderColor: '#ccc',
    borderWidth: 1,
  },
  row: {
    flexDirection: 'row',
    paddingVertical: 10,
    borderBottomWidth: 1,
    borderBottomColor: '#ccc',
  },
  cell: {
    flex: 1,
    textAlign: 'center',
    fontSize: 16,
  },
});
